// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ISignalService {
    struct Hop {
        address signalRootRelay;
        bytes32 signalRoot;
        bytes storageProof;
    }

    struct Proof {
        address crossChainSync;
        uint64 height;
        bytes storageProof;
        Hop[] hops;
    }

    function proveSignalReceived(
        uint64 srcChainId,
        address app,
        bytes32 signal,
        bytes calldata proof
    )
        external
        view
        returns (bool);
    
    function _foo(
        uint64 srcChainId,
        address app,
        bytes32 signal,
        Proof calldata proof
    )
        external
        view
        returns (bool);

    function sendSignal(bytes32 signal) external returns (bytes32 storageSlot);

    function getSignalSlot(
        uint64 chainId,
        address app,
        bytes32 signal
    ) external pure returns (bytes32 signalSlot);

    /// @notice Verifies if a particular signal has already been sent.
    /// @param app The address that initiated the signal.
    /// @param signal The signal (message) to send.
    /// @return True if the signal has been sent, otherwise false.
    function isSignalSent(address app, bytes32 signal) external view returns (bool);
}

abstract contract SignalService is ISignalService {}
