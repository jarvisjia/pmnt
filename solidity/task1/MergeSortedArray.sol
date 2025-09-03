// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract MergeSortedArray {

    //将两个有序数组合并为一个有序数组。
    function merge(uint256[] memory nums1, uint256[] memory nums2) public pure returns (uint256[] memory) {
        uint256[] memory nums = new uint256[](nums1.length + nums2.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 c = 0;
        while (i < nums1.length && j < nums2.length) {
            if (nums1[i] < nums2[j]) {
                nums[c] = nums1[i];
                i++;
            } else {
                nums[c] = nums2[j];
                j++;
            }
            c++;
        }
        while (i < nums1.length) {
            nums[c] = nums1[i];
            i++;
            c++;
        }
        while (j < nums2.length) {
            nums[c] = nums2[i];
            j++;
            c++;
        }
        return nums;
    }
}