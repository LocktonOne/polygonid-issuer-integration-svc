// SPDX-License-Identifier: MIT
pragma solidity 0.8.16;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {GenesisUtils} from "../lib/GenesisUtils.sol";
import {ICircuitValidator} from "../interfaces/ICircuitValidator.sol";
import {IZKPVerifier} from "../interfaces/IZKPVerifier.sol";
import {ArrayUtils} from "../lib/ArrayUtils.sol";

contract ZKPVerifier is IZKPVerifier, Ownable {
    /**
     * @dev Max return array length for request queries
     */
    uint256 public constant REQUESTS_RETURN_LIMIT = 1000;

    /// @dev Main storage structure for the contract
    struct MainStorage {
        // msg.sender-> ( requestID -> is proof given )
        mapping(address => mapping(uint64 => bool)) proofs;
        mapping(uint64 => IZKPVerifier.ZKPRequest) _requests;
        uint64[] _requestIds;
    }

    // keccak256(abi.encode(uint256(keccak256("iden3.storage.ZKPVerifier")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 internal constant ZKP_VERIFIER_VERIFIER_STORAGE_LOCATION =
    0x512d18c55869273fec77e70d8a8586e3fb133e90f1db24c6bcf4ff3506ef6a00;

    /// @dev Get the main storage using assembly to ensure specific storage location
    function _getMainStorage() internal pure returns (MainStorage storage $) {
        assembly {
            $.slot := ZKP_VERIFIER_VERIFIER_STORAGE_LOCATION
        }
    }

    function submitZKPResponse(
        uint64 requestId,
        uint256[] calldata inputs,
        uint256[2] calldata a,
        uint256[2][2] calldata b,
        uint256[2] calldata c
    ) public override {
        MainStorage storage s = _getMainStorage();
        IZKPVerifier.ZKPRequest storage request = s._requests[requestId];

        require(
            request.validator != ICircuitValidator(address(0)),
            "validator is not set for this request id"
        ); // validator exists

        _beforeProofSubmit(requestId, inputs, request.validator);
        request.validator.verify(inputs, a, b, c, request.data, msg.sender);
        s.proofs[msg.sender][requestId] = true; // user provided a valid proof for request
        _afterProofSubmit(requestId, inputs, request.validator);
    }

    function getZKPRequest(
        uint64 requestId
    ) public view override returns (IZKPVerifier.ZKPRequest memory) {
        require(requestIdExists(requestId), "request id doesn't exist");
        return _getMainStorage()._requests[requestId];
    }

    function setZKPRequest(
        uint64 requestId,
        ZKPRequest calldata request
    ) public override onlyOwner {
        MainStorage storage s = _getMainStorage();
        s._requests[requestId] = request;
        s._requestIds.push(requestId);
    }

    function getZKPRequestsCount() public view returns (uint256) {
        return _getMainStorage()._requestIds.length;
    }

    function requestIdExists(uint64 requestId) public view override returns (bool) {
        MainStorage storage s = _getMainStorage();
        for (uint i = 0; i < s._requestIds.length; i++) {
            if (s._requestIds[i] == requestId) {
                return true;
            }
        }

        return false;
    }

    function getZKPRequests(
        uint256 startIndex,
        uint256 length
    ) public view returns (IZKPVerifier.ZKPRequest[] memory) {
        MainStorage storage s = _getMainStorage();
        (uint256 start, uint256 end) = ArrayUtils.calculateBounds(
            s._requestIds.length,
            startIndex,
            length,
            REQUESTS_RETURN_LIMIT
        );

        IZKPVerifier.ZKPRequest[] memory result = new IZKPVerifier.ZKPRequest[](end - start);

        for (uint256 i = start; i < end; i++) {
            result[i - start] = s._requests[s._requestIds[i]];
        }

        return result;
    }

    /**
     * @dev Hook that is called before any proof response submit
     */
    function _beforeProofSubmit(
        uint64 requestId,
        uint256[] memory inputs,
        ICircuitValidator validator
    ) internal virtual {}

    /**
     * @dev Hook that is called after any proof response submit
     */
    function _afterProofSubmit(
        uint64 requestId,
        uint256[] memory inputs,
        ICircuitValidator validator
    ) internal virtual {}
}
