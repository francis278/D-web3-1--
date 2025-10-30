// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract BinarySearch {

    // 二分查找 (Binary Search)
    // 题目描述：在一个有序数组中查找目标值。

    function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256) {

        uint256 left = 0;
        uint256 right = arr.length - 1;

        // 循环条件
        while(left <= right){
            // 找到中间位
             uint256 mid = (left + right) / 2;

            // 中间值小于目标
            if (arr[mid] < target){
                left = mid + 1;
                // 中间值大于目标
            } else if (arr[mid] > target) {
                right = mid-1;
                // 中间值等于
            } else {
                return int256(mid);
            }

        }
        return -1;
    }
}