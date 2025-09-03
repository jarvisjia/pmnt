// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract BinarySearch {
    
    //在有序数组中查找目标值，如果找到返回下标，否则返回-1。
    function search(uint256[] memory nums, uint256 target) public pure returns (int256) {
        int256 left = 0;
        int256 right = int256(nums.length) - 1;
        while (left <= right) {
            int256 mid = (right + left) / 2;
            if (nums[uint256(mid)] == target) {
                return mid;
            } else if (nums[uint256(mid)] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
}