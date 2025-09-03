// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract ReverseString {

    //反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverseString(string memory str) public pure returns (string memory) {
        bytes memory bStr = bytes(str);
        for (uint i = 0; i < bStr.length / 2; i++) {
            (bStr[i], bStr[bStr.length - i - 1]) = (bStr[bStr.length - i - 1], bStr[i]);
        }
        return string(bStr);
    }
}