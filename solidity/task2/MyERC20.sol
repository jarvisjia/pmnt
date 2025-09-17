// SPDX-License-Identifier: MIT
pragma solidity ~0.8;

/*
ERC20 代币
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
    1.合约包含以下标准 ERC20 功能：
    2.balanceOf：查询账户余额。
    3.transfer：转账。
    4.approve 和 transferFrom：授权和代扣转账。
    5.使用 event 记录转账和授权操作。
    6.提供 mint 函数，允许合约所有者增发代币。
提示：
    使用 mapping 存储账户余额和授权信息。
    使用 event 定义 Transfer 和 Approval 事件。
    部署到sepolia 测试网，导入到自己的钱包
*/

contract MyERC20 {
    string private _name = "MyERC20";
    string private _symbol = "MTET";
    uint256 private _totalSupply;
    uint256 private _decimals = 18;
    mapping(address => uint256) private _balanceOf;
    mapping(address => mapping(address => uint256)) private allowance;
    address private owner;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /*
    检查权限
    */
    modifier onlyOwner() {
        require(owner == msg.sender, "Only owner can use this method");
        _;
    }

    constructor(uint256 _totalSup) {
        owner = msg.sender;
        _totalSupply = _totalSup * (10 ** uint256(_decimals));
        _balanceOf[owner] = _totalSupply;
        emit Transfer(address(0), owner, _totalSupply);
    }

    /*
    货币名称
    */
    function name() public view returns (string memory) {
        return _name;
    }

    /*
    货币符号
    */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /*
    货币精度单位
    */
    function decimals() public view returns (uint256) {
        return _decimals;
    }

    /*
    查询账户余额
    */
    function balanceOf(address addr) public view returns (uint256) {
        return _balanceOf[addr];
    }

    /*
    货币总供应量
    */
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    /*
    转账
    */
    function transfer(address to, uint256 value) public returns (bool) {
        require(to != address(0), "invalid address");
        require(value <= _balanceOf[msg.sender], "Insufficient balance");
        _balanceOf[msg.sender] -= value;
        _balanceOf[to] += value;
        emit Transfer(msg.sender, to, value);
        return true;
    }

    /*
    授权
    */
    function approve(address spender, uint256 value) public returns (bool) {
        require(spender != address(0), "invalid address");
        allowance[msg.sender][spender] = value;
        emit Approval(msg.sender, spender, value);
        return true;
    }

    /*
    代扣转账
    */
    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        require(to != address(0), "invalid address");
        require(value <= _balanceOf[from], "Insufficient balance");
        require(value <= allowance[from][msg.sender], "Insufficient allowance");
        _balanceOf[from] -= value;
        _balanceOf[to] += value;
        allowance[from][msg.sender] -= value;
        emit Transfer(from, to, value);
        return true;
    }
    
    /*
    铸币
    */
    function mint(address account, uint256 value) public onlyOwner returns (bool)  {
        require(account != address(0), "invalid address");
        uint256 valueWithDecimals = value * (10 ** uint256(_decimals));
        _totalSupply += valueWithDecimals;
        _balanceOf[account] += valueWithDecimals;
        emit Transfer(address(0), account, valueWithDecimals);
        return true;
    }
}