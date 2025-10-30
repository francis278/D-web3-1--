// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {
    // 定义符号结构体，存储罗马符号和对应的数值
    struct Symbol {
        string symbol;
        uint256 value;
    }

    // 符号映射数组，按数值从大到小排序（包含减法形式）
    Symbol[] private symbols;

    // 有效范围常量
    uint256 public constant MIN_VALUE = 1;
    uint256 public constant MAX_VALUE = 3999;

    // 构造函数：初始化所有罗马数字符号
    constructor() {
        // 按数值从大到小顺序初始化，包含所有基本符号和减法形式
        symbols.push(Symbol("M", 1000));
        symbols.push(Symbol("CM", 900));  // 特殊减法形式：1000-100
        symbols.push(Symbol("D", 500));
        symbols.push(Symbol("CD", 400));  // 特殊减法形式：500-100
        symbols.push(Symbol("C", 100));
        symbols.push(Symbol("XC", 90));   // 特殊减法形式：100-10
        symbols.push(Symbol("L", 50));
        symbols.push(Symbol("XL", 40));   // 特殊减法形式：50-10
        symbols.push(Symbol("X", 10));
        symbols.push(Symbol("IX", 9));    // 特殊减法形式：10-1
        symbols.push(Symbol("V", 5));
        symbols.push(Symbol("IV", 4));    // 特殊减法形式：5-1
        symbols.push(Symbol("I", 1));
    }

    /**
     * @dev 将整数转换为罗马数字
     * @param number 要转换的整数（1-3999）
     * @return 罗马数字字符串
     */
    function toRoman(uint256 number) public view returns (string memory) {
        // 输入验证：检查数字是否在有效范围内
        require(number >= MIN_VALUE && number <= MAX_VALUE, "Number must be between 1 and 3999");

        // 初始化结果字符串 
        string memory result = "";
        uint256 remaining = number;
        
        // 遍历符号数组，从最大值开始
        for (uint256 i = 0; i < symbols.length; i++) {
            Symbol memory currentSymbol = symbols[i];

            // 计算当前符号可以出现的次数
            uint256 count = remaining / currentSymbol.value;

            // 如果当前符号可以出现，则添加到结果中
            if (count > 0) {
                // 将符号重复count次并拼接到结果中
                result = string(abi.encodePacked(
                    result, 
                    _repeatSymbol(currentSymbol.symbol, count)
                ));

                // 减去已处理的部分
                remaining -= count * currentSymbol.value;
            }

            // 如果剩余值为0，提前结束循环
            if (remaining == 0){
                break;
            }
        }
        return result;
    }

    /**
     * @dev 内部函数：将符号重复指定次数
     * @param symbol 要重复的符号
     * @param count 重复次数
     * @return 重复后的字符串
     */
    function _repeatSymbol(string memory symbol, uint256 count) internal pure returns(string memory) {
        // 初始化结果字符串
        string memory result = "";

        // 循环重复符号
        for (uint256 i = 0; i < count; i++) {
            result = string(abi.encodePacked(result, symbol));
        }
        return result;
    }

    /**
     * @dev 批量转换多个数字为罗马数字
     * @param numbers 要转换的数字数组
     * @return 罗马数字字符串数组
     */
    function toRomanBatch(uint256[] memory numbers) public view returns (string[] memory) {
        // 创建结果数组
        string[] memory results = new string[](numbers.length);
        // 遍历数组并转换每个数字
        for (uint256 i = 0; i < numbers.length; i++) {
            results[i] = toRoman(numbers[i]);
        }
        return results;
    }

    /**
     * @dev 获取符号映射信息（用于调试和验证）
     * @return 所有符号的数组
     */
    function getSymbols() public view returns (Symbol[] memory) {
        return symbols;
    }

    /**
     * @dev 验证数字是否在有效范围内
     * @param number 要验证的数字
     * @return 是否有效
     */
    function isValidNumber(uint256 number) public pure returns (bool) {
        return number >= MIN_VALUE && number <= MAX_VALUE;
    }

    /**
     * @dev 获取转换范围信息
     * @return minValue 最小值
     * @return maxValue 最大值
     */
    function getRange() public pure returns (uint256 minValue, uint256 maxValue) {
        return (MIN_VALUE, MAX_VALUE);
    }
}