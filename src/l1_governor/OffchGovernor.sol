// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IGovernor, Governor} from "@openzeppelin/contracts/governance/Governor.sol";
import {GovernorVotes} from "@openzeppelin/contracts/governance/extensions/GovernorVotes.sol";
import {GovernorVotesQuorumFraction} from "@openzeppelin/contracts/governance/extensions/GovernorVotesQuorumFraction.sol";
import {Time} from "@openzeppelin/contracts/utils/types/Time.sol";
import {GovernorCountingFromSignal} from "./GovernorCountingFromSignal.sol";
import {ISignalService} from "../thirdparty/SignalService/SignalService.sol";
import {IVotes} from "@openzeppelin/contracts/governance/utils/IVotes.sol";

contract OffchGovernor is
    Governor,
    GovernorCountingFromSignal,
    GovernorVotes,
    GovernorVotesQuorumFraction
{
    ISignalService private signalService;
    uint64 private messageID = 0;
    address private votingToken;

    uint48 private snapshotBlockNumber = 0;

    // Define an event to emit after a message is sent
    event MessageSent(bytes32 messageHash, Message message);

    constructor(
        IVotes _votingToken,
        address _signalService
    )
        Governor("OffchGovernor")
        GovernorCountingFromSignal(_signalService)
        GovernorVotes(_votingToken)
        GovernorVotesQuorumFraction(15)
    {
        votingToken = address(_votingToken);
        signalService = ISignalService(_signalService);
    }

    function votingDelay() public pure override returns (uint256) {
        return 0;
    }

    function votingPeriod() public pure override returns (uint256) {
        // return 300; // 1 hour
        return 20; // 4 minutes
    }

    function proposalThreshold() public pure override returns (uint256) {
        return 0;
    }

    // Struct representing a message sent to L2.
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
        uint256 snapshotBlockNumber;
        uint256 duration;
        address votingTokenAddress;
    }

    // This functions overrides required by Offch Technology
    function _castVote(
        uint256 proposalId,
        address account,
        uint8 support,
        string memory reason,
        bytes memory params
    ) internal override(Governor) returns (uint256) {
        revert("OffchGoverner: you have to cast vote through L2");
    }

    // TODO: remove this function after information about l1 blocks that used to propose block to l2
    function proposeWithSnapshotBlockNumber(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        string memory description,
        uint48 _snapshotBlockNumber
    ) public virtual returns (uint256) {
        snapshotBlockNumber = _snapshotBlockNumber;

        address proposer = _msgSender();

        // check description restriction
        if (!_isValidDescriptionForProposer(proposer, description)) {
            revert GovernorRestrictedProposer(proposer);
        }

        // check proposal threshold
        uint256 proposerVotes = getVotes(proposer, clock() - 1);
        uint256 votesThreshold = proposalThreshold();
        if (proposerVotes < votesThreshold) {
            revert GovernorInsufficientProposerVotes(
                proposer,
                proposerVotes,
                votesThreshold
            );
        }

        uint256 proposalID = _propose(
            targets,
            values,
            calldatas,
            description,
            proposer
        );

        snapshotBlockNumber = 0;

        return proposalID;
    }

    function _propose(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        string memory description,
        address proposer
    ) internal override(Governor) returns (uint256 proposalId) {
        proposalId = super._propose(
            targets,
            values,
            calldatas,
            description,
            proposer
        );
 
        uint256 _snapshotBlockNumber = block.number;
        if (snapshotBlockNumber !=0) {
            _snapshotBlockNumber = snapshotBlockNumber;
        }

        Message memory message = Message({
            id: messageID,
            from: address(this),
            srcChainId: 17000,
            destChainId: 167008,
            proposalId: proposalId,
            snapshotBlockNumber: _snapshotBlockNumber,
            duration: votingPeriod(),
            votingTokenAddress: votingToken
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

    // This functions overrides required by Solidity

    function clock()
        public
        view
        override(Governor, GovernorVotes)
        returns (uint48)
    {
        // TODO: block number will be retrieved from taiko L1 smart contract from the map that contains block number of L1 block that used to propose block to L2
        if (snapshotBlockNumber > 0) {
            return snapshotBlockNumber;
        }

        return Time.blockNumber();
    }

    // function CLOCK_MODE() public view override returns (string memory) {
    //     return "mode=blocknumber&from=default";
    // }
}
