// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInteger {
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
        symbols.push(Symbol("D", 500));
        symbols.push(Symbol("C", 100));
        symbols.push(Symbol("L", 50));
        symbols.push(Symbol("X", 10));
        symbols.push(Symbol("V", 5));
        symbols.push(Symbol("I", 1));
    }

    /**
     * @dev 将罗马数字转换为整数
     * @param roman 要转换的罗马数字字符串
     * @return 转换后的整数值
     */
    function toInteger(string memory roman) public view returns (uint256) {
        // 将输入字符串转换为字节数组以便处理
        bytes memory romanBytes = bytes(roman);
        // 检查输入字符串是否为空
        require(romanBytes.length > 0, "Input string cannot be empty");

        // 初始化结果为0
        uint256 result = 0;

        // 遍历罗马数字的每个字符（从第一个到倒数第二个）
        for (uint256 i = 0; i < romanBytes.length - 1; i++) {
            // 获取当前字符的数值
            uint256 currentValue = _charToValue(romanBytes[i]);
            // 获取下一个字符的数值
            uint256 nextValue = _charToValue(romanBytes[i + 1]);

            // 如果当前值小于下一个值，执行减法规则
            if (currentValue < nextValue) {
                result -= currentValue; // 从结果中减去当前值
            } else {
                result += currentValue; // 否则将当前值加到结果中
            }
        }

        // 处理最后一个字符（总是加到结果中）
        result += _charToValue(romanBytes[romanBytes.length - 1]);

        // 验证结果是否在有效范围内
        require(result >= MIN_VALUE && result <= MAX_VALUE, "Invalid Roman numeral");

        // 返回最终的计算结果
        return result;
    }

    /**
     * @dev 内部函数：将单个罗马数字字符转换为对应的数值
     * @param c 罗马数字字符
     * @return 对应的整数值
     */
    function _charToValue(bytes1 c) internal view returns (uint256) {
        // 遍历符号数组查找匹配的字符
        for (uint256 i = 0; i < symbols.length; i++) {
            // 将字符串符号转换为字节进行比较
            bytes memory symbolBytes = bytes(symbols[i].symbol);

            // 检查是否匹配单个字符的符号
            if (symbolBytes.length == 1 && symbolBytes[0] == c) {
                return symbols[i].value;
            }
        }
        revert("Invalid Roman numeral character"); // 如果字符无效，回滚交易
    }

    /**
     * @dev 内部函数：检查字符是否有效
     * @param c 罗马数字字符
     * @return 是否有效
     */
    function _isValidChar(bytes1 c) internal view returns(bool){
        // 遍历符号数组查找匹配的字符
        for (uint256 i = 0; i< symbols.length; i++) {
            // 将字符串符号转换为字节进行比较
            bytes memory symbolBytes = bytes(symbols[i].symbol);

            // 检查是否匹配单个字符的符号
            if (symbolBytes.length == 1 && symbolBytes[0] == c){
                return true;
            }
        }
        return false;
    }

    /**
     * @dev 验证罗马数字字符串是否有效
     * @param roman 要验证的罗马数字字符串
     * @return 是否有效
     */
    function isValidRoman(string memory roman) public view returns (bool) {
        // 将字符串转换为字节数组
        bytes memory romanBytes = bytes(roman);

        // 检查字符串是否为空
        if (romanBytes.length == 0) {
            return false;
        }

        // 检查每个字符是否都是有效的罗马数字字符
        for (uint256 i = 0; i < romanBytes.length; i++) {
        if (!_isValidChar(romanBytes[i])) {
            return false;
        }
        }

        // 所有字符都有效，返回true
        return true;
    }

    /**
     * @dev 安全地将罗马数字转换为整数，包含有效性检查
     * @param roman 罗马数字字符串
     * @return 转换后的整数值
     */
    function safeToInteger(string memory roman) public view returns (uint256) {
        // 检查字符串是否有效
        require(isValidRoman(roman), "Invalid Roman numeral");
        // 如果有效，进行转换
        return toInteger(roman);
    }

    /**
     * @dev 批量转换多个罗马数字为整数
     * @param romans 要转换的罗马数字字符串数组
     * @return 整数数组
     */
    function toIntegerBatch(string[] memory romans) public view returns (uint256[] memory) {
        // 创建结果数组
        uint256[] memory results = new uint256[](romans.length);

        // 遍历数组并转换每个罗马数字
        for (uint256 i = 0; i < romans.length; i++) {
            results[i] = toInteger(romans[i]);
        }

        // 返回转换结果数组
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
     * @dev 获取转换范围信息
     * @return minValue 最小值
     * @return maxValue 最大值
     */
    function getRange() public pure returns (uint256 minValue, uint256 maxValue) {
        return (MIN_VALUE, MAX_VALUE);
    }
}