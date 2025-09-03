// SPDX-License-Identifier: MIT
pragma solidity ~0.8.0;

contract IntToRoman {

     struct Roman {
        string symbol;
        uint256 value;
    }

    Roman[] private romanNumerals;

    constructor() {
        romanNumerals.push(Roman("M", 1000));
        romanNumerals.push(Roman("CM", 900));
        romanNumerals.push(Roman("D", 500));
        romanNumerals.push(Roman("CD", 400));
        romanNumerals.push(Roman("C", 100));
        romanNumerals.push(Roman("XC", 90));
        romanNumerals.push(Roman("L", 50));
        romanNumerals.push(Roman("XL", 40));
        romanNumerals.push(Roman("X", 10));
        romanNumerals.push(Roman("IX", 9));
        romanNumerals.push(Roman("V", 5));
        romanNumerals.push(Roman("IV", 4));
        romanNumerals.push(Roman("I", 1));
    }

    function intToRoman(uint256 num) public view returns (string memory) {
        require(num > 0 && num < 4000, "Input must be between 1 and 3999");
        string memory result = "";
        for (uint i = 0; i < romanNumerals.length; i++) {
            while (num >= romanNumerals[i].value) {
                result = string(abi.encodePacked(result, romanNumerals[i].symbol));
                num -= romanNumerals[i].value;
            }
        }
        return result;
    }
}