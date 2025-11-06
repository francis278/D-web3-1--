// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BeggingContract {

// 作业3：编写一个讨饭合约
// 任务目标
// 使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
// 记录每个捐赠者的地址和捐赠金额。
// 允许合约所有者提取所有捐赠的资金。

// 任务步骤
// 编写合约
// 创建一个名为 BeggingContract 的合约。
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。

// 部署合约
// 在 Remix IDE 中编译合约。
// 部署合约到 Goerli 或 Sepolia 测试网。
// 测试合约
// 使用 MetaMask 向合约发送以太币，测试 donate 功能。
// 调用 withdraw 函数，测试合约所有者是否可以提取资金。
// 调用 getDonation 函数，查询某个地址的捐赠金额。

// 任务要求
// 合约代码：
// 使用 mapping 记录捐赠者的地址和金额。
// 使用 payable 修饰符实现 donate 和 withdraw 函数。
// 使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
// 测试网部署：
// 合约必须部署到 Goerli 或 Sepolia 测试网。
// 功能测试：
// 确保 donate、withdraw 和 getDonation 函数正常工作。

// 提交内容
// 合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。
// 合约地址：提交部署到测试网的合约地址。
// 测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。

// 额外挑战（可选）
// 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
// 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
// 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。



    mapping( address => uint256 ) public donations;

    // 存储前3名捐赠者
    address[3] public topDonors;
    uint256[3] public topAmounts;

    address public owner;

    constructor() {
        owner = msg.sender;
    }

    event DonationMade(address indexed from, address indexed to, uint256 value); 
    event PaymentMade(address indexed from, address indexed to, uint256 value);
    event Withdrawal(address indexed to, uint256 value);

    modifier onlyOwner() {
        require(owner == msg.sender, "Not the contract owner");
        _;
    }

    modifier onlyDuringDonationHours() {
        require(isWithinDonationHours(), "Donations only allowed from 10:00 to 16:00");
        _;
    }

    function isWithinDonationHours() public view returns (bool) {
        // 获取当前时间（UTC）
        uint256 currentHour = (block.timestamp / 3600) % 24;
        
        // 检查是否在 10:00-16:00 之间（UTC时间）
        return currentHour >= 10 && currentHour < 16;
        // 检查是否在 00:00-01:00 之间（UTC时间）
        //return currentHour >= 0 && currentHour < 1;
    }

    // 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
    function donate(address to, uint256 amount) public payable onlyDuringDonationHours returns (bool) {
        require(msg.value >= amount, "Insufficient funds");
        require(amount > 0, "Donation amount must be greater than 0");
        require(to != address(0), "Cannot donate to zero address");
        donations[to] += amount;

        // 更新排行榜
        updateLeaderboard(msg.sender, amount);
        emit DonationMade(msg.sender, to, amount);
        emit Donation(msg.sender, amount);
        return true;
    }

    // 一个 withdraw 函数，允许合约所有者提取所有资金。
    function withdraw(address to) public payable onlyOwner {
        require(owner == msg.sender, "Not the contract owner");
        require(to != address(0), "Cannot withdraw to zero address");
        uint256 balance = address(this).balance;
        payable(to).transfer(balance);
        emit Withdrawal(to, balance);
    }

    // 一个 getDonation 函数，允许查询某个地址的捐赠金额。
    function getDonation(address addr) public view returns (uint256){
        return donations[addr];
    }

    // 使用 payable 修饰符和 address.transfer 实现支付
    function Payment(address to, uint256 amount) public payable returns (bool){
        require(msg.value == amount, "Sent ETH must match amount");
        require(amount > 0, "Payment amount must be greater than 0");
        require(to != address(0), "Cannot pay to zero address");

        payable(to).transfer(amount);
        emit PaymentMade(msg.sender, to, amount);
        return true;
    }

    // 使用 payable 修饰符和 address.transfer 实现提款。
    function withdrawal(address to, uint256 amount) public returns (bool){
        require(owner == msg.sender, "Not the contract owner");
        require(to != address(0), "Cannot donate to zero address");
        require(address(this).balance >= amount, "Insufficient funds");

        payable(to).transfer(amount);
        emit Withdrawal(to, amount);
        return true;
    }

    // 额外挑战（可选）
    // 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
     event Donation(address indexed donor, uint256 amount);
    // 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
    // 每次捐赠时更新排行榜
    function updateLeaderboard(address donor, uint256 amount) internal {

        // 如果新金额小于等于第3名且第3名不为0，直接返回
        if (topAmounts[2] != 0 && amount <= topAmounts[2]) {
            return;
        }
        
        // 遍历找到新金额应该插入的位置
        int256 insertPosition = -1;
        
        for (int256 i = 2; i >= 0; i--) {
            uint256 index = uint256(i);
            
            // 如果当前位置为空或者新金额更大，记录插入位置
            if (topAmounts[index] == 0 || amount > topAmounts[index]) {
                insertPosition = i;
            } else {
                // 一旦遇到比新金额大的，就停止搜索
                break;
            }
        }
        
        // 如果没有找到插入位置，说明金额太小
        if (insertPosition == -1) {
            return;
        }
        
        // 移动元素，为插入腾出空间
        for (uint256 j = 2; j > uint256(insertPosition); j--) {
            topDonors[j] = topDonors[j - 1];
            topAmounts[j] = topAmounts[j - 1];
        }
        
        // 插入新元素
        topDonors[uint256(insertPosition)] = donor;
        topAmounts[uint256(insertPosition)] = amount;
    }

    // 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
    function getTopDonors() public view returns (address[3] memory, uint256[3] memory){
        return (topDonors, topAmounts);
    }
    
    function getContractBalance() public view returns (uint256) {
        return address(this).balance;
    }


}