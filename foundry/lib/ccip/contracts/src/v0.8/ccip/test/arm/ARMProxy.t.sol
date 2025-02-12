// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IARM} from "../../interfaces/IARM.sol";

import {ARM} from "../../ARM.sol";
import {ARMProxy} from "../../ARMProxy.sol";
import {MockARM} from "../mocks/MockARM.sol";
import {ARMSetup} from "./ARMSetup.t.sol";

contract ARMProxyTest is ARMSetup {
  event ARMSet(address arm);

  ARMProxy internal s_armProxy;

  function setUp() public virtual override {
    ARMSetup.setUp();
    s_armProxy = new ARMProxy(address(s_arm));
  }

  function test_ARMIsCursed_Success() public {
    s_armProxy.setARM(address(s_mockARM));
    assertFalse(IARM(address(s_armProxy)).isCursed());
    ARM(address(s_armProxy)).voteToCurse(bytes32(0));
    assertTrue(IARM(address(s_armProxy)).isCursed());
  }

  function test_ARMIsBlessed_Success() public {
    s_armProxy.setARM(address(s_mockARM));
    assertTrue(IARM(address(s_armProxy)).isBlessed(IARM.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
    ARM(address(s_armProxy)).voteToCurse(bytes32(0));
    assertFalse(IARM(address(s_armProxy)).isBlessed(IARM.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
  }

  function test_ARMCallRevertReasonForwarded() public {
    bytes memory err = bytes("revert");
    s_mockARM.setRevert(err);
    s_armProxy.setARM(address(s_mockARM));
    vm.expectRevert(abi.encodeWithSelector(MockARM.CustomError.selector, err));
    IARM(address(s_armProxy)).isCursed();
  }
}
