// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";



abstract contract Group is Initializable, UUPSUpgradeable {

    struct Group {
        string name;
        mapping(address => bool) members;
    }

    function initialize() initializer public {
        __UUPSUpgradeable_init();
    }


}


contract Group is Initializable, UUPSUpgradeable {

    function initialize() initializer public {
        __UUPSUpgradeable_init();
    }


}