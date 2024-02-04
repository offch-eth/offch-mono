// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {OffchGovernor} from "../l1_governor/OffchGovernor.sol";
import {BlockHeader, LibBlockHeader} from "../libs/LibBlockHeader.sol";
import {LibSecureMerkleTrie} from "../thirdparty/LibSecureMerkleTrie.sol";
import {LibRLPReader} from "../thirdparty/LibRLPReader.sol";
import {CrossChainSync} from "../thirdparty/CrossChainSync/CrossChainSync.sol";
import {ISignalService} from "../thirdparty/SignalService/SignalService.sol";

contract ProposalVoteFactory {
    ISignalService private signalService;
    CrossChainSync private crossChainSync;
    uint64 private srcChainId;
    uint64 private messageID = 0;

    enum VoteType {
        Against,
        For,
        Abstain
    }

    struct ProposalVote {
        address app;
        uint256 proposalId;
        uint256 snapshotBlockNumber;
        uint256 duration;
        address votingTokenAddress;
        uint256 againstVotes;
        uint256 forVotes;
        uint256 abstainVotes;
        mapping(address voter => bool) hasVoted;
    }

    mapping(uint256 key => ProposalVote) private _proposalVotes;

    // Define an event to emit after a message is sent
    event MessageSent(bytes32 messageHash, Message message);

    // Struct representing a message sent to L1.
    struct Message {
        // Message ID.
        uint64 id;
        // Message sender address (auto filled).
        address from;
        // Source chain ID (auto filled).
        uint64 srcChainId;
        // Destination chain ID where the `to` address lives (auto filled).
        uint64 destChainId;
        uint256 proposalId;
        uint256 againstVotes;
        uint256 forVotes;
        uint256 abstainVotes;
    }

    constructor(
        uint64 _srcChainId,
        address _signalService,
        address _crossChainSync
    ) {
        srcChainId = _srcChainId;
        signalService = ISignalService(_signalService);
        crossChainSync = CrossChainSync(_crossChainSync);
    }

    event ProposalCreated(
        address indexed app,
        address indexed votingTokenAddress,
        uint256 snapshotBlockNumber,
        uint256 duration,
        bytes32 messageHash,
        uint256 voteKey
    );

    event VoteCast(
        uint256 indexed voteKey,
        address indexed voter,
        uint8 support,
        uint256 weight
    );

    function createProposalVoteFromSignal(
        address app,
        bytes32 _messageHash,
        OffchGovernor.Message memory message,
        bytes calldata proof
    ) public {
        require(
            signalService.proveSignalReceived(
                srcChainId,
                app,
                _messageHash,
                proof
            ),
            "Invalid signal."
        );

        bytes memory bytesMessage = abi.encode(message);

        // Hash the message to fit the bytes32 signal requirement
        bytes32 messageHash = keccak256(bytesMessage);

        // Verify that the message hash is the same as the one passed in the signal
        require(
            messageHash == _messageHash,
            "Message is different from the signal."
        );

        uint256 voteKey = uint256(
            keccak256(abi.encodePacked(messageHash, app))
        );

        ProposalVote storage proposalVote = _proposalVotes[voteKey];

        proposalVote.app = app;
        proposalVote.snapshotBlockNumber = message.snapshotBlockNumber;
        proposalVote.duration = message.duration;
        proposalVote.votingTokenAddress = message.votingTokenAddress;
        proposalVote.proposalId = message.proposalId;

        emit ProposalCreated(
            app,
            message.votingTokenAddress,
            message.snapshotBlockNumber,
            message.duration,
            messageHash,
            voteKey
        );
    }

    function sendSignalAboutResults(uint256 key) public {
        ProposalVote storage vote = _proposalVotes[key];

        // Check if the vote duration is over
        require(
            crossChainSync.latestSyncedL1Height() >=
                vote.snapshotBlockNumber + vote.duration,
            "Voting period is not over yet"
        );

        Message memory message = Message({
            id: messageID,
            from: address(this),
            srcChainId: 167008,
            destChainId: 17000,
            proposalId: vote.proposalId,
            forVotes: vote.forVotes,
            againstVotes: vote.againstVotes,
            abstainVotes: vote.abstainVotes
        });

        messageID++;

        // Combine the two strings into a single bytes array
        bytes memory bytesMessage = abi.encode(message);

        // Hash the combined message to fit the bytes32 signal requirement
        bytes32 messageHash = keccak256(bytesMessage);

        // Send the message hash as a signal to the Signal Service
        signalService.sendSignal(messageHash);

        // Emit the event with the message hash and the original strings
        emit MessageSent(messageHash, message);
    }

    event SignalSent(uint256 indexed key, bool result);

    function hasVoted(
        uint256 key,
        address account
    ) public view virtual returns (bool) {
        return _proposalVotes[key].hasVoted[account];
    }

    function proposalVotes(
        uint256 key
    )
        public
        view
        virtual
        returns (uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
    {
        ProposalVote storage proposalVote = _proposalVotes[key];

        return (
            proposalVote.againstVotes,
            proposalVote.forVotes,
            proposalVote.abstainVotes
        );
    }

    struct ProofsToVerify {
        bytes tokenAddress;
        bytes tokenAccountState;
        bytes tokenAccountProof;
        bytes32 stateRoot;
        bytes storageKey;
        bytes tokenAmount;
        bytes storageProof;
        bytes32 storageRootHash;
    }

    struct BlockHeaderToRebuild {
        uint64 blockNumber;
        BlockHeader blockHeader;
    }

    function countVote(
        uint256 key,
        uint8 support,
        ProofsToVerify calldata content,
        BlockHeaderToRebuild memory blockHeader
    ) public {
        ProposalVote storage proposalVote = _proposalVotes[key];
        require(
            proposalVote.snapshotBlockNumber > 0,
            "Proposal vote does not exist."
        );

        // Check if the vote duration is over
        require(
            crossChainSync.latestSyncedL1Height() <
                proposalVote.snapshotBlockNumber + proposalVote.duration,
            "Voting period is over"
        );

        address voter = msg.sender;

        require(
            proposalVote.votingTokenAddress ==
                address(bytes20(content.tokenAddress)),
            "Invalid voting token address."
        );

        require(
            proposalVote.hasVoted[voter] == false,
            "Account has already voted on this proposal."
        );

        _verifyTokenOwnage(content, blockHeader);

        uint256 weight = LibRLPReader.readUint256(content.tokenAmount);

        proposalVote.hasVoted[voter] = true;

        if (support == uint8(VoteType.Against)) {
            proposalVote.againstVotes += weight;
        } else if (support == uint8(VoteType.For)) {
            proposalVote.forVotes += weight;
        } else if (support == uint8(VoteType.Abstain)) {
            proposalVote.abstainVotes += weight;
        } else {
            revert("Invalid vote type.");
        }

        emit VoteCast(key, voter, support, weight);
    }

    function _verifyTokenOwnage(
        ProofsToVerify calldata content,
        BlockHeaderToRebuild memory blockHeader
    ) internal view virtual returns (bool) {
        CrossChainSync.Snippet memory snippet = crossChainSync.getSyncedSnippet(
            blockHeader.blockNumber
        );

        bool blockHashValid = false;

        bool accountProofValid = _verifyProof(
            content.tokenAddress,
            content.tokenAccountState,
            content.tokenAccountProof,
            content.stateRoot
        );
        bool storageProofValid = _verifyProof(
            content.storageKey,
            content.tokenAmount,
            content.storageProof,
            content.storageRootHash
        );

        if (accountProofValid && storageProofValid) {
            blockHeader.blockHeader.stateRoot = content.stateRoot;

            blockHashValid =
                snippet.blockHash ==
                LibBlockHeader.hashBlockHeader(blockHeader.blockHeader);
        }

        if (!blockHashValid) {
            revert(
                string(
                    abi.encodePacked(
                        "Proof is wrong based on the block hash: ",
                        snippet.blockHash
                    )
                )
            );
        }

        return true;
    }

    function _verifyProof(
        bytes calldata key,
        bytes calldata value,
        bytes calldata proof,
        bytes32 root
    ) private pure returns (bool valid) {
        return
            LibSecureMerkleTrie.verifyInclusionProof(key, value, proof, root);
    }
}
