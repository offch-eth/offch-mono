// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Governor} from "@openzeppelin/contracts/governance/Governor.sol";
import {ISignalService} from "../thirdparty/SignalService/SignalService.sol";

// import {ProposalVoteFactory} from "../l2_voting/ProposalVote.sol";

abstract contract GovernorCountingFromSignal is Governor {
    ISignalService private signalService;

    /**
     * @dev Supported vote types. Matches Governor Bravo ordering.
     */
    enum VoteType {
        Against,
        For,
        Abstain
    }

    struct ProposalVote {
        uint256 againstVotes;
        uint256 forVotes;
        uint256 abstainVotes;
        bool isSet;
    }

    mapping(uint256 proposalId => ProposalVote) private _proposalVotes;

    struct ProposalVoteMessage {
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

    constructor(address _signalService) {
        signalService = ISignalService(_signalService);
    }

    function resultFromSignal(
        address app,
        bytes32 _messageHash,
        ProposalVoteMessage memory message,
        bytes calldata proof
    ) public {
        require(
            signalService.proveSignalReceived(
                message.srcChainId,
                app,
                _messageHash,
                proof
            ),
            "Invalid signal."
        );

        require(
            !_proposalVotes[message.proposalId].isSet,
            "Results are already set for this proposal." 
        );

        ProposalVote storage proposalVote = _proposalVotes[message.proposalId];
        proposalVote.againstVotes = message.againstVotes;
        proposalVote.forVotes = message.forVotes;
        proposalVote.abstainVotes = message.abstainVotes;
        proposalVote.isSet = true;
    }

    /**
     * @dev See {IGovernor-COUNTING_MODE}.
     */
    // solhint-disable-next-line func-name-mixedcase
    function COUNTING_MODE()
        public
        pure
        virtual
        override
        returns (string memory)
    {
        return "support=bravo&quorum=for,abstain";
    }

    /**
     * @dev See {IGovernor-hasVoted}.
     */
    function hasVoted(
        uint256 proposalId,
        address account
    ) public view virtual override returns (bool) {
        // TODO: if for each vote there will be signal from L2 to L1
        // It can be verified if account has voted
        return false;
    }

    /**
     * @dev Accessor to the internal vote counts.
     */
    function proposalVotes(
        uint256 proposalId
    )
        public
        view
        virtual
        returns (uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
    {
        ProposalVote storage proposalVote = _proposalVotes[proposalId];
        return (
            proposalVote.againstVotes,
            proposalVote.forVotes,
            proposalVote.abstainVotes
        );
    }

    /**
     * @dev See {Governor-_quorumReached}.
     */
    function _quorumReached(
        uint256 proposalId
    ) internal view virtual override returns (bool) {
        ProposalVote storage proposalVote = _proposalVotes[proposalId];

        return
            quorum(proposalSnapshot(proposalId)) <=
            proposalVote.forVotes + proposalVote.abstainVotes;
    }

    /**
     * @dev See {Governor-_voteSucceeded}. In this module, the forVotes must be strictly over the againstVotes.
     */
    function _voteSucceeded(
        uint256 proposalId
    ) internal view virtual override returns (bool) {
        ProposalVote storage proposalVote = _proposalVotes[proposalId];

        return proposalVote.forVotes > proposalVote.againstVotes;
    }

    /**
     * @dev See {Governor-_countVote}. In this module, the support follows the `VoteType` enum (from Governor Bravo).
     */
    function _countVote(
        uint256 proposalId,
        address account,
        uint8 support,
        uint256 weight,
        bytes memory // params
    ) internal virtual override {
        // raise error that you can vote only through L2
        revert("GovernorCountingFromSignal: you have to cast vote through L2");
    }
}
