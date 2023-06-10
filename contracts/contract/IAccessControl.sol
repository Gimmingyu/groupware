// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IAccessControl {

    struct RoleData {
        mapping(address => bool) members;
        bytes32 adminRole;
    }

    event RoleAdminChanged();

    event RoleGranted();

    event RoleRevoked();
}