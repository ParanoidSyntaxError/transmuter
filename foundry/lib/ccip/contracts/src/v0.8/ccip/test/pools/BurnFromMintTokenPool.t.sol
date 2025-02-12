// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Pool} from "../../libraries/Pool.sol";
import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";
import {BurnFromMintTokenPool} from "../../pools/BurnFromMintTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";

contract BurnFromMintTokenPoolSetup is BurnMintSetup {
  BurnFromMintTokenPool internal s_pool;

  function setUp() public virtual override {
    BurnMintSetup.setUp();

    s_pool = new BurnFromMintTokenPool(s_burnMintERC677, new address[](0), address(s_mockARM), address(s_sourceRouter));
    s_burnMintERC677.grantMintAndBurnRoles(address(s_pool));

    _applyChainUpdates(address(s_pool));
  }
}

contract BurnFromMintTokenPool_lockOrBurn is BurnFromMintTokenPoolSetup {
  function test_Setup_Success() public view {
    assertEq(address(s_burnMintERC677), address(s_pool.getToken()));
    assertEq(address(s_mockARM), s_pool.getArmProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq(type(uint256).max, s_burnMintERC677.allowance(address(s_pool), address(s_pool)));
    assertEq("BurnFromMintTokenPool 1.5.0-dev", s_pool.typeAndVersion());
  }

  function test_PoolBurn_Success() public {
    uint256 burnAmount = 20_000e18;

    deal(address(s_burnMintERC677), address(s_pool), burnAmount);
    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), burnAmount);

    vm.startPrank(s_burnMintOnRamp);

    vm.expectEmit();
    emit TokensConsumed(burnAmount);

    vm.expectEmit();
    emit Transfer(address(s_pool), address(0), burnAmount);

    vm.expectEmit();
    emit Burned(address(s_burnMintOnRamp), burnAmount);

    bytes4 expectedSignature = bytes4(keccak256("burnFrom(address,uint256)"));
    vm.expectCall(address(s_burnMintERC677), abi.encodeWithSelector(expectedSignature, address(s_pool), burnAmount));

    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: burnAmount,
        remoteChainSelector: DEST_CHAIN_SELECTOR
      })
    );

    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), 0);
  }

  // Should not burn tokens if cursed.
  function test_PoolBurnRevertNotHealthy_Revert() public {
    s_mockARM.voteToCurse(bytes32(0));
    uint256 before = s_burnMintERC677.balanceOf(address(s_pool));
    vm.startPrank(s_burnMintOnRamp);

    vm.expectRevert(EVM2EVMOnRamp.BadARMSignal.selector);
    s_pool.lockOrBurn(
      Pool.LockOrBurnInV1({
        originalSender: OWNER,
        receiver: bytes(""),
        amount: 1e5,
        remoteChainSelector: DEST_CHAIN_SELECTOR
      })
    );

    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), before);
  }

  function test_ChainNotAllowed_Revert() public {
    uint64 wrongChainSelector = 8838833;
    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
    s_pool.releaseOrMint(
      Pool.ReleaseOrMintInV1({
        originalSender: bytes(""),
        receiver: OWNER,
        amount: 1,
        remoteChainSelector: wrongChainSelector,
        sourcePoolAddress: generateSourceTokenData().sourcePoolAddress,
        sourcePoolData: generateSourceTokenData().extraData,
        offchainTokenData: ""
      })
    );
  }
}
