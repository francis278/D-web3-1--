https://sepolia.etherscan.io/address/0x2761e723a57f1436dd5259fe0a9eb40cad71cc60



[block:9564257 txIndex:11]from: 0x8a2...f6533to: MyNFT.mintNFT(address,string) 0x276...1cc60value: 0 weidata: 0xeac...00000logs: 2hash: 0xe1a...34047
status	0x1 Transaction mined and execution succeed
transaction hash	0x82da6c82033eb6822f7cbfad4603f0656296dcce1eaf01a1459f6eb59c99406e
block hash	0xe1acd5ee79274cdd2797a613f8a07c37b1eeb6b17407aaffbc1fb33ac4534047
block number	9564257
from	0x8A20b414387D602F0665EA6D84060c9b81FF6533
to	MyNFT.mintNFT(address,string) 0x2761e723a57f1436dd5259fe0a9eb40cad71cc60
gas	164742 gas
transaction cost	162287 gas 
input	0xeac...00000
decoded input	{
	"address recipient": "0x8A20b414387D602F0665EA6D84060c9b81FF6533",
	"string tokenURI": "bafkreif4sn3aqwhzpqva5z324npxnpbc7rmjcw7qlkfhn4fonyzufayb7q"
}
decoded output	 - 
logs	[
	{
		"from": "0x2761e723a57f1436dd5259fe0a9eb40cad71cc60",
		"topic": "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		"event": "Transfer",
		"args": {
			"0": "0x0000000000000000000000000000000000000000",
			"1": "0x8A20b414387D602F0665EA6D84060c9b81FF6533",
			"2": "1"
		}
	},
	{
		"from": "0x2761e723a57f1436dd5259fe0a9eb40cad71cc60",
		"topic": "0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7",
		"event": "MetadataUpdate",
		"args": {
			"0": "1"
		}
	}
]
raw logs	[
  {
    "address": "0x2761e723a57f1436dd5259fe0a9eb40cad71cc60",
    "blockHash": "0xe1acd5ee79274cdd2797a613f8a07c37b1eeb6b17407aaffbc1fb33ac4534047",
    "blockNumber": "0x91f061",
    "data": "0x",
    "logIndex": "0x18",
    "removed": false,
    "topics": [
      "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
      "0x0000000000000000000000000000000000000000000000000000000000000000",
      "0x0000000000000000000000008a20b414387d602f0665ea6d84060c9b81ff6533",
      "0x0000000000000000000000000000000000000000000000000000000000000001"
    ],
    "transactionHash": "0x82da6c82033eb6822f7cbfad4603f0656296dcce1eaf01a1459f6eb59c99406e",
    "transactionIndex": "0xb"
  },
  {
    "address": "0x2761e723a57f1436dd5259fe0a9eb40cad71cc60",
    "blockHash": "0xe1acd5ee79274cdd2797a613f8a07c37b1eeb6b17407aaffbc1fb33ac4534047",
    "blockNumber": "0x91f061",
    "data": "0x0000000000000000000000000000000000000000000000000000000000000001",
    "logIndex": "0x19",
    "removed": false,
    "topics": [
      "0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7"
    ],
    "transactionHash": "0x82da6c82033eb6822f7cbfad4603f0656296dcce1eaf01a1459f6eb59c99406e",
    "transactionIndex": "0xb"
  }
]


测试合约
使用 MetaMask 向合约发送以太币，测试 donate 功能。
[block:9566273 txIndex:9]from: 0x8a2...f6533to: BeggingContract.donate(address,uint256) 0x816...0dfb3value: 1000000000000000000 weidata: 0xe69...40000logs: 1hash: 0x4ba...f1b8c
status	0x1 Transaction mined and execution succeed
transaction hash	0x3101df30b4925e73050a1b83c4ceeeba1d6fa8cab50fde0884cab0d2554d3059
block hash	0x4ba8cf70ae754f0f64b73ae4417de95a5af9cae0099d91e04c414ccf217f1b8c
block number	9566273
from	0x8A20b414387D602F0665EA6D84060c9b81FF6533
to	BeggingContract.donate(address,uint256) 0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3
gas	47303 gas
transaction cost	46919 gas 
input	0xe69...40000
decoded input	{
	"address to": "0x8167C4dB4c639761acBcd0f2BFc590332cD0DFB3",
	"uint256 amount": "1000000000000000000"
}
decoded output	 - 
logs	[
	{
		"from": "0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3",
		"topic": "0x70c6bc6cf53707fb5561feef795d557f7e46f57bd4f1a13ce36e9ffdf58d38bb",
		"event": "DonationMade",
		"args": {
			"0": "0x8A20b414387D602F0665EA6D84060c9b81FF6533",
			"1": "0x8167C4dB4c639761acBcd0f2BFc590332cD0DFB3",
			"2": "1000000000000000000"
		}
	}
]
raw logs	[
  {
    "address": "0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3",
    "blockHash": "0x4ba8cf70ae754f0f64b73ae4417de95a5af9cae0099d91e04c414ccf217f1b8c",
    "blockNumber": "0x91f841",
    "data": "0x0000000000000000000000000000000000000000000000000de0b6b3a7640000",
    "logIndex": "0xd",
    "removed": false,
    "topics": [
      "0x70c6bc6cf53707fb5561feef795d557f7e46f57bd4f1a13ce36e9ffdf58d38bb",
      "0x0000000000000000000000008a20b414387d602f0665ea6d84060c9b81ff6533",
      "0x0000000000000000000000008167c4db4c639761acbcd0f2bfc590332cd0dfb3"
    ],
    "transactionHash": "0x3101df30b4925e73050a1b83c4ceeeba1d6fa8cab50fde0884cab0d2554d3059",
    "transactionIndex": "0x9"
  }
]
value	1000000000000000000 wei



调用 withdraw 函数，测试合约所有者是否可以提取资金。
[block:9566357 txIndex:7]from: 0x8a2...f6533to: BeggingContract.withdraw(address) 0x816...0dfb3value: 0 weidata: 0x51c...f6533logs: 1hash: 0xb60...980d6
status	0x1 Transaction mined and execution succeed
transaction hash	0x81cb26e253bc05032961b986649ee70eee4c066cd496a9f52c3025738a6867ef
block hash	0xb60985b93412eac5d0660b3042f554aa8b59bcbdd6df43ba9ae7b941778980d6
block number	9566357
from	0x8A20b414387D602F0665EA6D84060c9b81FF6533
to	BeggingContract.withdraw(address) 0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3
gas	33308 gas
transaction cost	32596 gas 
input	0x51c...f6533
decoded input	{
	"address to": "0x8A20b414387D602F0665EA6D84060c9b81FF6533"
}
decoded output	 - 
logs	[
	{
		"from": "0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3",
		"topic": "0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65",
		"event": "Withdrawal",
		"args": {
			"0": "0x8A20b414387D602F0665EA6D84060c9b81FF6533",
			"1": "1000000000000000000"
		}
	}
]
raw logs	[
  {
    "address": "0x8167c4db4c639761acbcd0f2bfc590332cd0dfb3",
    "blockHash": "0xb60985b93412eac5d0660b3042f554aa8b59bcbdd6df43ba9ae7b941778980d6",
    "blockNumber": "0x91f895",
    "data": "0x0000000000000000000000000000000000000000000000000de0b6b3a7640000",
    "logIndex": "0xd",
    "removed": false,
    "topics": [
      "0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65",
      "0x0000000000000000000000008a20b414387d602f0665ea6d84060c9b81ff6533"
    ],
    "transactionHash": "0x81cb26e253bc05032961b986649ee70eee4c066cd496a9f52c3025738a6867ef",
    "transactionIndex": "0x7"
  }
]

