// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Reverse {

    string public originalStr;

    // 反转字符串 (Reverse String)
    // 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverseString(string memory inputStr) public pure returns (string memory) {
        bytes memory originalBytes = bytes(inputStr);
        uint256 inputLength = originalBytes.length;
        bytes memory reversedBytes = new bytes(inputLength);
        for (uint256 i = 0; i < inputLength; i++) 
        {
            reversedBytes[inputLength-1-i] =originalBytes[i];
        }

        return string(reversedBytes);
    }
}