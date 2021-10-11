pragma solidity ^0.5.1;
import "./nodes.sol";
import "./safemath.sol";


contract HpbVote is Ownable {
    using SafeMath for uint256;

    HpbNodes boenodes;
    address payable nodeaddr;

    function setNodeContract(address payable addr) onlyAdmin public {
        boenodes = HpbNodes(addr);
        nodeaddr = addr;
    }

    function getNodeContract() public view returns(address){
        return nodeaddr;
    }

    uint _gasLeftLimit=1000000;//对于过于复杂操作，无法一步完成，那么必须分步进行
    
    uint public minLimit=1 ether;//最小投票数量限额：1HPB,可外部设置
    
    struct BoeIndex{
        address coinbase;
        uint index;
        uint voteNumber;
    }
    
    struct Voter{
        address payable voterAddr;//投票人
        uint voteNumber;//投的总票数
        address[] boes;
        mapping (address=> BoeIndex) vote; //投票明细 boe =>boeindex
    }
    Voter[] voterArray;//投票者的数组
    
    struct VoterIndex{
        address coinbase;
        uint index;
    }
    mapping (address => VoterIndex) voterIndexMap;//投票者地址=>投票者数组下标
    
    // struct VoteInfo{
    //     address voter;
    //     uint voteNumber;
    // }
    // mapping (address => VoteInfo[]) voterinfos;
    mapping (address => uint) voteresult;//boe => num
    
    /**
     * 获取某个候选人被投票详细情况
     */
    function fetchVoteInfoForCandidate(address boeaddr) public view returns(address[] memory,uint[] memory){
        address[] memory voterinfo = new address[](voterArray.length);
        uint[] memory votenum = new uint[](voterArray.length);
        uint p = 0;
        for(uint i = 0; i < voterArray.length; i++){
            if(voterArray[i].vote[boeaddr].voteNumber > 0 ){
                voterinfo[p] = (voterArray[i].voterAddr);
                votenum[p] = voterArray[i].vote[boeaddr].voteNumber;
                p++;
            }
        }
        address[] memory resinfo = new address[](p);
        uint[] memory resnum = new uint[](p);
        for (uint k = 0;k < p; k++){
            resinfo[k] = voterinfo[k];
            resnum[k] = votenum[k];
        }
        return (resinfo,resnum);
    }

    /**
    *查询所有的节点票数
    */
    function fetchAllVoteResult() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        address payable[] memory nodes = boenodes.getAlllockNode();
        uint i = nodes.length;
        uint[] memory _nums=new uint[](i);
        for (uint k = 0; k < i; k++) {
            _nums[k] = voteresult[nodes[k]];
        }
        return (nodes,_nums);
    }
    
     function refreshVoteForAll() public {
        for (uint i=1;i<voterArray.length;i++) {
            if ( voterArray[i].voteNumber > 0 ) {
                if (voterArray[i].voterAddr.balance < voterArray[i].voteNumber){
                    docancelVote(voterArray[i].voterAddr);
        	    }
            }
            //这里的分步操作需要一次交易才会继续执行，谁来发交易？
            //因此存在失效的投票无法被清除的情况。谁来清除？
            //刷新的数据量会越来越多，导致投票越晚，需要消耗的gas越多
            //越晚的无效投票越不容易被清除
            if(gasleft()<_gasLeftLimit){ 
                break;
            }
        }
    }
    
    /**
     * 投票
     */
    function vote(address payable boeaddr,uint num) public{
        dovote(boeaddr,num);
        refreshVoteForAll();
    } 
     
    function dovote(
        address payable boeaddr,
        uint num
    )  internal {
        require(num >= minLimit,"num too low");
        require(boenodes.isLockNode(boeaddr),"unlock node");
        if (voterIndexMap[msg.sender].coinbase == address(0)){
            voterIndexMap[msg.sender].coinbase = msg.sender;
            voterIndexMap[msg.sender].index = voterArray.length;
            Voter memory newvoter;
            BoeIndex memory newboeindex;
            newvoter.voterAddr = msg.sender;
            newvoter.voteNumber = num;
            newboeindex.index = 0;
            newboeindex.voteNumber = num;
            newboeindex.coinbase = boeaddr;
            voterArray.push(newvoter);
            voterArray[voterArray.length-1].boes.push(boeaddr);
            voterArray[voterArray.length-1].vote[boeaddr] = newboeindex;
        }else{
            voterArray[voterIndexMap[msg.sender].index].voteNumber = num.add(voterArray[voterIndexMap[msg.sender].index].voteNumber);
            if (voterArray[voterIndexMap[msg.sender].index].vote[boeaddr].coinbase == address(0)){//之前未给boeddr投过票
                voterArray[voterIndexMap[msg.sender].index].vote[boeaddr].index = voterArray[voterIndexMap[msg.sender].index].boes.length;
                voterArray[voterIndexMap[msg.sender].index].vote[boeaddr].coinbase = boeaddr;
                voterArray[voterIndexMap[msg.sender].index].boes.push(boeaddr);
            }
            voterArray[voterIndexMap[msg.sender].index].vote[boeaddr].voteNumber = num.add(voterArray[voterIndexMap[msg.sender].index].vote[boeaddr].voteNumber);
        }
        voteresult[boeaddr] = num.add(voteresult[boeaddr]);
        require(msg.sender.balance >= voterArray[voterIndexMap[msg.sender].index].voteNumber,"balance too low");
    }


    /**
     * 用于批量投票 
     */
    function  batchvote(
        address payable[] memory boeaddrs,
        uint[] memory nums
    )  public {
        require(boeaddrs.length==nums.length);
        for(uint i = 0; i < boeaddrs.length;i++){
            dovote(boeaddrs[i],nums[i]);
        }
        refreshVoteForAll();
    }
     /*
    *撤销所有投票
    */
    function cancelVote() public{
        docancelVote(msg.sender);
    }
    
    function docancelVote(address voteraddr) internal{
        VoterIndex memory newindex = voterIndexMap[voteraddr];
        require(newindex.coinbase == voteraddr);
        uint i = newindex.index;
        for (uint j = 0; j < voterArray[i].boes.length; j++){ //撤销投票
            voteresult[voterArray[i].boes[j]] = voteresult[voterArray[i].boes[j]].sub(voterArray[i].vote[voterArray[i].boes[j]].voteNumber);
    	    voterArray[i].vote[voterArray[i].boes[j]].voteNumber = 0;
    	}
	    voterArray[i].voteNumber = 0;
    }
    
    /**
     * 撤回对某个候选人的投票 
     */
    function cancelVoteForCandidate(
        address payable boeaddr
    ) public returns(bool) {
        VoterIndex memory newindex = voterIndexMap[msg.sender];
        require(newindex.coinbase == msg.sender);
        uint i = newindex.index;
        for (uint j = 0; j < voterArray[i].boes.length; j++){ //撤销投票
            if (voterArray[i].boes[j] == boeaddr){
                voteresult[voterArray[i].boes[j]] = voteresult[voterArray[i].boes[j]].sub(voterArray[i].vote[voterArray[i].boes[j]].voteNumber);
    	        voterArray[i].voteNumber = voterArray[i].voteNumber.sub(voterArray[i].vote[voterArray[i].boes[j]].voteNumber);
    	        voterArray[i].vote[voterArray[i].boes[j]].voteNumber = 0;
            }
    	}
    }
    
    /**
     * 获取投票人的投票情况
     */
    function fetchVoteInfoForVoter(
        address voterAddr
    ) public view returns (
        address[] memory, 
        uint[] memory
    ) {
        VoterIndex memory newindex = voterIndexMap[voterAddr];
        if (newindex.coinbase == address(0)) { //没投过票 
            return (new address[](0), new uint[](0));
        }
        
        if (voterArray[newindex.index].voteNumber == 0) {
            return (new address[](0), new uint[](0));
        }
        uint i = voterArray[newindex.index].boes.length;
        address[] memory _addrs=new address[](i);
        uint[] memory _nums=new uint[](i);
        for (uint k = 0;k < i;k++) {
            _addrs[k] = voterArray[newindex.index].boes[k];
            _nums[k] = voterArray[newindex.index].vote[voterArray[newindex.index].boes[k]].voteNumber;
        }
        return (_addrs, _nums);
    }
    
     /**
     * 获取所有投票人的详细信息
     */
    function fetchAllVoters() public view returns (
        address payable[] memory, 
        uint[] memory
    ) {
        uint i = 0;
        for (uint j = 0;j < voterArray.length;j++) {
            if(voterArray[j].voteNumber > 0){
                i++;
            }
        }
        address payable[] memory _addrs=new address payable[](i);
        uint[] memory _voteNumbers=new uint[](i);
        uint p=0;
        for (uint k = 0;k < voterArray.length;k++) {
            if(voterArray[k].voteNumber > 0){
                _addrs[p]=voterArray[k].voterAddr;
                _voteNumbers[p] = voterArray[k].voteNumber;
                p++;
            }
        }
        return (_addrs, _voteNumbers);
    }
    
     /**
     * 获取某个候选人的总得票数 Total number of votes obtained from candidates
     */
    function fetchVoteNumForCandidate(
        address payable boeaddr
    ) public view returns (uint) {
        return voteresult[boeaddr];
    }
    
    /**
     * 获取某个投票人已投票数 Total number of votes obtained from voterAddr
     */
    function fetchVoteNumForVoter(
        address payable voterAddr
    ) public view returns (uint) {
        VoterIndex memory newindex = voterIndexMap[voterAddr];
        if (newindex.coinbase == voterAddr){
            return voterArray[newindex.index].voteNumber;
        }
        return 0;
    }
    
    /*
    *查询voterAddr给boeAddr的投票数
    */
    function fetchVoteNumForVoterToCandidate(
        address payable voterAddr,
        address payable boeAddr
    ) public view returns (uint) {
        VoterIndex memory newindex = voterIndexMap[voterAddr];
        if (newindex.coinbase == voterAddr){
            for (uint i = 0; i < voterArray[newindex.index].boes.length;i++){
                if (voterArray[newindex.index].boes[i] == boeAddr){
                    return voterArray[newindex.index].vote[boeAddr].voteNumber;
                }
            }
        }
        return 0;
    }
    
    
    
    //接受HPB转账
    function () payable external {
        emit ReceivedHpb(msg.sender, msg.value);
    }
    //销毁合约，并把合约余额返回给合约拥有者
    function kill() onlyOwner public returns(bool) {
        selfdestruct(owner);
        return true;
    }
    //合约拥有者提取合约余额的一部分
    // function withdraw(
    //     uint _value
    // ) onlyOwner payable public returns(bool) {
    //     require(address(this).balance >= _value);
    //     owner.transfer(_value);
    //     return true;
    // }
    
    event ApprovalFor(
        bool indexed approved,
        address payable indexed operator, 
        address payable indexed owner
    );
    
    event DoVoted(// 投票  flag=1为投票,flag=0为撤票
        uint indexed flag,
        address payable indexed candidateAddr,
        address payable indexed voteAddr,
        uint num
    );
    //记录发送HPB的发送者地址和发送的金额
    event ReceivedHpb(
        address payable indexed sender, 
        uint amount
    );
    
    /**
     *设置最小允许的投票数
     */
    function setMinLimit(
        uint _minLimit
    ) onlyAdmin public returns(bool) {
        require(_minLimit > 1 ether);
        minLimit = _minLimit;
        return true;
    }
    
    /**
     * 构造函数 初始化投票智能合约的部分依赖参数
     */
    constructor () payable public {
        owner = msg.sender;
        // 设置默认管理员
        adminMap[owner] = owner;
    }
    
    /**
     * 为了防止因为gas消耗超出而导致本次操作失败
     */
	function setGasLeftLimit(uint gasLeftLimit) onlyAdmin public returns (bool) {
        _gasLeftLimit=gasLeftLimit;
        return true;
    }
    
}