pragma solidity ^0.5.1;
import "./nodes.sol";
import "./lock.sol";
import "./vote.sol";
contract Proxy is Ownable{
    address node;
    address lock;
    address vote;
    HpbNodes hpbnode;
    HpbLock hpblock;
    HpbVote hpbvote;

    constructor () payable public {
        owner = msg.sender;
        addAdmin(owner);
    }

    function setnodecontract(address payable nodeaddr) onlyAdmin public{
        node = nodeaddr;
        hpbnode = HpbNodes(nodeaddr);
    }

    function setvotecontract(address payable voteaddr) onlyAdmin public{
        vote = voteaddr;
        hpbvote = HpbVote(voteaddr);
    }

    function setlockcontract(address payable lockaddr) onlyAdmin public{
        lock = lockaddr;
        hpblock = HpbLock(lockaddr);
    }

    function getcontract() public view returns(address,address,address){
        return (node,lock,vote);
    }

    function getAllHpbNodes(
    )  public view returns (
        address payable[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ){
        return hpbnode.getAllHpbNodes();
    }

    function fetchAllVoteResult() public view returns (
        address payable[] memory, 
        uint[] memory
    ){
        return hpbvote.fetchAllVoteResult();
    }

    function fetchAllHolderAddrs() public view returns(
        address[] memory,
        address[] memory){
        return hpbnode.fetchAllHolderAddrs();
    }

}