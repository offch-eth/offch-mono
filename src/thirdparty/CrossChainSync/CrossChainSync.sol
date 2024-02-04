// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ICrossChainSync {
    struct Snippet {
        uint64 remoteBlockId;
        uint64 syncedInBlock;
        bytes32 blockHash;
        bytes32 signalRoot;
    }
    
    /// @notice Fetches the hash of a block from the opposite chain.
    /// @param blockId The target block id. Specifying 0 retrieves the hash
    /// of the latest block.
    /// @return snippet The block hash and signal root synced.
    function getSyncedSnippet(uint64 blockId) external view returns (Snippet memory snippet);
}

abstract contract CrossChainSync is ICrossChainSync {
    uint64 public latestSyncedL1Height;
}
