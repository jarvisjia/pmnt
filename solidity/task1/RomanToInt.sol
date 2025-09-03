// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract RomanToInt {

    mapping(bytes1 => uint256) private romanMap;

    constructor() {
        romanMap["I"] = 1;
        romanMap["V"] = 5;
        romanMap["X"] = 10;
        romanMap["L"] = 50;
        romanMap["C"] = 100;
        romanMap["D"] = 500;
        romanMap["M"] = 1000;
    }

    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory strBytes = bytes(s);
        uint256 total = 0;
        uint256 length = strBytes.length;
        for (uint i = 0; i < length; i++) {
            uint256 currentValue = romanMap[strBytes[i]];
            if (i < length-1 &&  currentValue < romanMap[strBytes[i + 1]]) {
                total -= currentValue;
            } else {
                total += currentValue;
            }

        }
        return total;
    }
}