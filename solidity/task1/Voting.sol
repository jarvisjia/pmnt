// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

contract Voting {
    //存储候选人的得票数
    mapping(address => uint256) public votesReceived;
    //存储候选人
    address[] public receivers;

    //允许用户投票给某个候选人
    function vote(address _address) public {
        if (votesReceived[_address] == 0) {
            receivers.push(_address);
        }
        votesReceived[_address] += 1;
    }

    //返回某个候选人的得票数
    function getVotes(address _address) public view returns (uint256) {
        return votesReceived[_address];
    }

    //重置所有候选人的得票数
    function resetVotes() public {
        for (uint i = 0; i < receivers.length; i++) {
            votesReceived[receivers[i]] = 0;
        }
        receivers = new address[](0);
    }
}