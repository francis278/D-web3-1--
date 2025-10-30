// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortedArray {

    // 合并两个有序数组 (Merge Sorted Array)
    // 题目描述：将两个有序数组合并为一个有序数组。
    function mergeSortedArray(uint[] memory nums1, uint[] memory nums2) public pure returns (uint[] memory) {
        uint n1 = nums1.length;
        uint n2 = nums2.length;
        uint[] memory merged = new uint[](n1 + n2);
        uint i = 0; // nums1的索引
        uint j = 0; // nums2的索引
        uint k = 0; // 合并后数

        while (i < n1 && j < n2) {
            if (nums1[i] <= nums2[j]){
                merged[k] = nums1[i];
                i++;
            } else {
                merged[k] = nums2[j];
                j++;
            }
            k++;
        }

        while (i < n1) {
            merged[k] = nums1[i];
            i++;
            k++;
        }
        
        while (j < n2) {
            merged[k] = nums2[j];
            j++;
            k++;
        }

    return merged;
    }
}