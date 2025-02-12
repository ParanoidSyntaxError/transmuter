// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";
import {ITokenMessenger} from "../../pools/USDC/ITokenMessenger.sol";

// This contract mocks both the ITokenMessenger and IMessageTransmitter
// contracts involved with the Cross Chain Token Protocol.
contract MockUSDCTokenMessenger is ITokenMessenger {
  uint32 private immutable i_messageBodyVersion;
  bytes32 public constant i_destinationTokenMessenger = keccak256("i_destinationTokenMessenger");
  uint64 public s_nonce;
  address private i_transmitter;

  constructor(uint32 version, address transmitter) {
    i_messageBodyVersion = version;
    s_nonce = 1;
    i_transmitter = transmitter;
  }

  function depositForBurnWithCaller(
    uint256 amount,
    uint32 destinationDomain,
    bytes32 mintRecipient,
    address burnToken,
    bytes32 destinationCaller
  ) external returns (uint64) {
    IBurnMintERC20(burnToken).transferFrom(msg.sender, address(this), amount);
    IBurnMintERC20(burnToken).burn(amount);
    emit DepositForBurn(
      s_nonce,
      burnToken,
      amount,
      msg.sender,
      mintRecipient,
      destinationDomain,
      i_destinationTokenMessenger,
      destinationCaller
    );
    return s_nonce++;
  }

  function messageBodyVersion() external view returns (uint32) {
    return i_messageBodyVersion;
  }

  function localMessageTransmitter() external view returns (address) {
    return i_transmitter;
  }
}
