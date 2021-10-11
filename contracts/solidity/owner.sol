pragma solidity ^0.5.1;

contract Ownable {
    address payable public owner;
    address payable[] admins;
    modifier onlyOwner {
        require(msg.sender == owner);
        // Do not forget the "_;"! It will be replaced by the actual function
        // body when the modifier is used.
        _;
    }

    function transferOwnership(
        address payable newOwner
    ) onlyOwner public returns(bool) {
        owner = newOwner;
        addAdmin(newOwner);
        deleteAdmin(owner);
        return true;
    }

    function getOwner() public view returns (
        address payable
    ) {
        return owner;
    }
    // 合约管理员，可以添加和删除候选人
    mapping (address  => address payable) public adminMap;

    modifier onlyAdmin {
        require(adminMap[msg.sender] != address(0));
        _;
    }

    function addAdmin(
        address payable addr
    ) onlyOwner public returns(bool) {
        require(adminMap[addr] == address(0));
        adminMap[addr] = addr;
        admins.push(addr);
        return true;
    }

    function deleteAdmin(
        address payable addr
    ) onlyOwner public returns(bool) {
        require(adminMap[addr] != address(0));
        adminMap[addr] = address(0);
        delete adminMap[addr];
        for(uint256 i = 0; i < admins.length; i++){
            if (admins[i] == addr){
                admins[i] = admins[admins.length-1];
                delete admins[admins.length-1];
                admins.length--;
            }
        }
        return true;
    }
    
    function getAdmins() public view returns(address payable[] memory){
        return admins;
    }
}