// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";



abstract contract DocumentStore is Initializable, UUPSUpgradeable {

    function initialize() initializer public {
        __UUPSUpgradeable_init();
    }


}


contract DocumentStore is Initializable, UUPSUpgradeable {

    function initialize() initializer public {
        __UUPSUpgradeable_init();
    }


}