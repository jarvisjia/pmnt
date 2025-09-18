// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

/*
编写一个讨饭合约
任务目标
1. 使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
2. 记录每个捐赠者的地址和捐赠金额。
3. 允许合约所有者提取所有捐赠的资金。

任务步骤
1. 编写合约
  - 创建一个名为 BeggingContract 的合约。
  - 合约应包含以下功能：
  - 一个 mapping 来记录每个捐赠者的捐赠金额。
  - 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
  - 一个 withdraw 函数，允许合约所有者提取所有资金。
  - 一个 getDonation 函数，允许查询某个地址的捐赠金额。
  - 使用 payable 修饰符和 address.transfer 实现支付和提款。
2. 部署合约
  - 在 Remix IDE 中编译合约。
  - 部署合约到 Goerli 或 Sepolia 测试网。
3. 测试合约
  - 使用 MetaMask 向合约发送以太币，测试 donate 功能。
  - 调用 withdraw 函数，测试合约所有者是否可以提取资金。
  - 调用 getDonation 函数，查询某个地址的捐赠金额。

任务要求
1. 合约代码：
  - 使用 mapping 记录捐赠者的地址和金额。
  - 使用 payable 修饰符实现 donate 和 withdraw 函数。
  - 使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
2. 测试网部署：
  - 合约必须部署到 Goerli 或 Sepolia 测试网。
3. 功能测试：
  - 确保 donate、withdraw 和 getDonation 函数正常工作。

提交内容
1. 合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。
2. 合约地址：提交部署到测试网的合约地址。
3. 测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。

额外挑战（可选）
1. 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
2. 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
3. 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
*/

import "@openzeppelin/contracts/access/Ownable.sol";

contract BeggingContract is Ownable {

    // 记录每个地址的捐赠总额
    mapping (address=>uint256) public donations;

    //捐赠排行版结构
    struct Donor {
        address donorAddress;
        uint256 amount;
    }
    //显示捐赠金额最多的前3个地址
    Donor[3] public topDonors;

    //时间限制
    uint256 public donationStartTime;
    //捐赠期为3天
    uint256 public constant DONATION_PERIOD = 3 days;

    // 捐赠事件
    event Donation(address indexed donor, uint256 amount);
    // 提取事件
    event Withdraw(address indexed to, uint256 amount);

    modifier onlyDuringDonationPeriod() {
        require(block.timestamp >= donationStartTime && block.timestamp <= donationStartTime + DONATION_PERIOD, "Donation period has ended");
        _;
    }

    // 构造函数，设置合约部署者为所有者
    constructor() Ownable(msg.sender) {
        //从部署时开始计算捐赠期
        donationStartTime = block.timestamp;
    }

    // 捐赠函数，允许用户向合约发送以太币
    function donate() external payable onlyDuringDonationPeriod {
        require(msg.value > 0, "value must be greater then zero");

        // 记录捐赠金额
        donations[msg.sender] += msg.value;

        // 更新排行榜
        _updateTopDonors(msg.sender);

        // 触发捐赠事件
        emit Donation(msg.sender, msg.value);
    }

    // 提取所有资金，仅所有者可调用
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        // 将所有余额转账给所有者
        payable(owner()).transfer(balance);

        // 触发提取事件
        emit Withdraw(owner(), balance);
    }

    // 查询某个地址的捐赠总额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    // 获取合约当前余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // 获取合约当前地址
    function getContractAddress() external view returns (address) {
        return address(this);
    }

    // 获取合约当前地址
    function getMsgSender() external view returns (address) {
        return msg.sender;
    }

    //获取捐赠剩余时间
    function getRemainingDonnatioinTime() external view returns (uint256) {
        if (block.timestamp>donationStartTime+DONATION_PERIOD) {
            return 0;
        }
        return (donationStartTime + DONATION_PERIOD) - block.timestamp;
    }

    //所有者可以延长捐赠时间
    function extendDonationPeriod(uint256 _additionalTime) external onlyOwner {
        donationStartTime += _additionalTime;
    }

     //更新排行版
    function _updateTopDonors(address _donor) internal {
        uint256 currentDonation = donations[_donor];
        //检查是否应该进入排行版
        for (uint8 i = 0; i<3; i++) {
            if (currentDonation>topDonors[i].amount) {
                // 将新的捐赠者插入到合适的位置
                for (uint8 j = 2; j>i; j--) {
                    topDonors[j] = topDonors[j-1];
                }
                topDonors[i] = Donor(_donor,currentDonation);
                break;
            }
        }
    }
}