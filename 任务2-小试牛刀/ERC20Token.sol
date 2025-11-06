// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ERC20Token {

    mapping(address=>uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;
    uint256 private totalSupply_;
    address public owner; 

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    // 构造函数中设置所有者
    constructor() {
        owner = msg.sender;
    }
    
    // 合约包含以下标准 ERC20 功能：
    function allowance(address account, address spender) public view returns (uint256) {
        return allowances[account][spender];
    }
    function totalSupply() public view returns (uint256) {
        return totalSupply_;  // 或者直接返回 totalSupply
    }

    // balanceOf：查询账户余额。
    function balanceOf(address account) public view returns (uint256){
        return balances[account];
    }

    // transfer：转账。
    function transfer(address to, uint256 amount) public returns (bool){

        require(balances[msg.sender] >= amount, "Insufficient balance");
        require(to != address(0), "ERC20: transfer to the zero address"); 

        balances[msg.sender] -= amount;
        balances[to] += amount;

        emit Transfer(msg.sender, to, amount);
        return true;
    }
    // approve 和 transferFrom：授权和代扣转账。
    function approve(address spender, uint256 value) public returns (bool){
        require(msg.sender != address(0), "ERC20: approve from zero address");
        require(spender != address(0), "ERC20: approve to zero address");
        allowances[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    function transferFrom(address from, address to, uint256 value) public returns (bool){
         // 1. 检查零地址
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");
        uint256 currentAllowance = allowances[from][msg.sender];
        require(currentAllowance >= value, "ERC20: insufficient allowance");
        require(balances[from] >= value, "ERC20: transfer amount exceeds balance");

        allowances[from][msg.sender] -= value;
        balances[from] -= value;
        balances[to] += value;

        emit Transfer(from, to, value);
        return true;
    }

    // 使用 event 记录转账和授权操作。
    // 提供 mint 函数，允许合约所有者增发代币。
    function mint(address to, uint256 value) public returns (bool){
        require(msg.sender == owner, "Only owner can call this function");
        require(to != address(0), "ERC20: mint to the zero address");
        totalSupply_ += value;
        balances[to] += value;
        emit Transfer(address(0), to, value);
        return true;
    }
} 