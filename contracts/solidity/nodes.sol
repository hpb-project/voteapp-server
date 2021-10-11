pragma solidity ^0.5.1;
import "./owner.sol";

/***
 * HPB节点信息
 *  */
contract HpbNodes is Ownable{
    string public name = "HPB Nodes Service";
	event ReceivedHpb(
        address payable indexed sender, 
        uint amount
    );
    event AddBoeNode(
        address indexed coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 indexed hid
	);
	event UpdateBoeNode(
        address indexed coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 indexed hid
	);
	event DeleteBoeNode(
	    address indexed coinbase
	);
	// 接受HPB转账
    function () payable external {
        emit ReceivedHpb(
            msg.sender, 
            msg.value
        );
    }
    // 销毁合约，并把合约余额返回给合约拥有者
    function kill() onlyOwner payable public returns(bool) {
        selfdestruct(owner);
        return true;
    }

    // function withdraw(
    //     uint _value
    // ) onlyOwner payable public returns(bool) {
    //     require(address(this).balance >= _value);
    //     owner.transfer(_value);
    //     return true;
    // }
    struct BoeNode{
        address payable coinbase;
        bytes32 cid1;
        bytes32 cid2;
        bytes32 hid;
    }
    
    struct NodeIndex{
        address coinbase;
        uint256 index;
    }
    
    BoeNode[] BoeNodes;
    mapping (address => NodeIndex) BoeNodesIndexMap;
    address[] LockNodes;
    mapping (address => NodeIndex) LockNodesIndexMap;
    
    constructor () payable public {
        owner = msg.sender;
        addAdmin(owner);
    }

    /**
	 * 添加节点信息
	 */
    function addBoeNode(
        address payable coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 hid
    ) onlyAdmin public returns(bool){
        require(BoeNodesIndexMap[coinbase].coinbase == address(0));
        BoeNode memory boe;
        boe.coinbase = coinbase;
        boe.cid1 = cid1;
        boe.cid2 = cid2;
        boe.hid = hid;
        BoeNodes.push(boe);
        BoeNodesIndexMap[coinbase].coinbase = coinbase;
        BoeNodesIndexMap[coinbase].index = BoeNodes.length-1;
        emit AddBoeNode(coinbase,cid1,cid2,hid);
        return true;
    }
    
    function addBoeNodeBatch(
        address payable[] memory coinbases,
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(addBoeNode(coinbases[i],cid1s[i],cid2s[i],hids[i]));
        }
        return true;
    }
	
    function updateBoeNode(
        address payable coinbase,
        bytes32 cid1,
        bytes32 cid2,
        bytes32 hid
    ) onlyAdmin public returns(bool){
        require(BoeNodesIndexMap[coinbase].coinbase != address(0));
        
        uint256 index = BoeNodesIndexMap[coinbase].index;
        BoeNodes[index].cid1 = cid1;
        BoeNodes[index].cid2 = cid2;
        BoeNodes[index].hid = hid;
        
        emit UpdateBoeNode(coinbase,cid1,cid2,hid);
        return true;
    }
    
    function updateBoeNodeBatch(
        address payable[] memory coinbases,
        bytes32[] memory cid1s,
        bytes32[] memory cid2s,
        bytes32[] memory hids
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(updateBoeNode(coinbases[i],cid1s[i],cid2s[i],hids[i]));
        }
        return true;
    }
    
    function deleteBoeNode(
        address payable coinbase
    ) onlyAdmin public returns(bool){
        //存在才需要删除
        require(BoeNodesIndexMap[coinbase].coinbase != address(0));
        
        //无质押才能删除
        require(LockNodesIndexMap[coinbase].coinbase == address(0));
        
        //删除boe的同时要删除质押索引，质押数据等到解除质押时在删除
        BoeNodes[BoeNodesIndexMap[coinbase].index] = BoeNodes[BoeNodes.length-1];
        BoeNodesIndexMap[BoeNodes[BoeNodes.length-1].coinbase].index = BoeNodesIndexMap[coinbase].index; 
        delete BoeNodes[BoeNodes.length-1];
        delete(BoeNodesIndexMap[coinbase]);
        BoeNodes.length--;
        emit DeleteBoeNode(coinbase);
        return true;
    }
    
    function deleteBoeNodeBatch(
        address payable[] memory coinbases
    ) onlyAdmin public returns(bool){
        for(uint i = 0;i<coinbases.length;i++){
            require(deleteBoeNode(coinbases[i]));
        }
        return true;
    }

    function getAllBoesAddrs() public view returns(address[] memory){
        //查询所有的Boe节点
        uint length=BoeNodes.length;
        address[] memory _coinbases=new address[](length);
        for (uint i = 0;i < length;i++){
            _coinbases[i]=BoeNodes[i].coinbase;
        }
        return _coinbases;
    }

    function getAllBoes() public view returns(
        address[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ){
        //查询所有的Boe节点
        uint length=BoeNodes.length;
        address[] memory _coinbases=new address[](length);
        bytes32[] memory _cid1s=new bytes32[](length);
        bytes32[] memory _cid2s=new bytes32[](length);
        bytes32[] memory _hids=new bytes32[](length);
        for (uint i = 0;i < length;i++){
            _coinbases[i]=BoeNodes[i].coinbase;
            _cid1s[i]=BoeNodes[i].cid1;
            _cid2s[i]=BoeNodes[i].cid2;
            _hids[i]=BoeNodes[i].hid;
        }
        return (_coinbases,_cid1s,_cid2s,_hids);
    }
   
    function getAllHpbNodes(
    )  public view returns (
        address payable[] memory,
        bytes32[] memory,
        bytes32[] memory,
        bytes32[] memory
    ){
        //只查询质押的节点
        uint length = LockNodes.length;
        address payable [] memory _coinbases = new address payable[](length);
        bytes32[] memory _cid1s = new bytes32[](length);
        bytes32[] memory _cid2s = new bytes32[](length);
        bytes32[] memory _hids = new bytes32[](length);
        for (uint i = 0;i < length;i++){
            uint256 index = BoeNodesIndexMap[LockNodes[i]].index;
            _coinbases[i] = BoeNodes[index].coinbase;
            _cid1s[i] = BoeNodes[index].cid1;
            _cid2s[i] = BoeNodes[index].cid2;
            _hids[i] = BoeNodes[index].hid;
        }
        return (_coinbases,_cid1s,_cid2s,_hids);
    }
    
    function getAlllockNode(
    )  public view returns (
        address payable[] memory
    ){
        //只查询质押的节点
        uint length = LockNodes.length;
        address payable[] memory _coinbases = new address payable[](length);
        for (uint i = 0;i < length;i++){
            uint256 index = BoeNodesIndexMap[LockNodes[i]].index;
            _coinbases[i] = BoeNodes[index].coinbase;
        }
        return _coinbases;
    }
    
    function isBoeNode(address addr) public view returns(bool){
        return BoeNodesIndexMap[addr].coinbase != address(0);
    }
    
    function isLockNode(address addr) public view returns(bool){
        return LockNodesIndexMap[addr].coinbase != address(0);
    }

    address LockContractAddr;
    
    function setLockContract(address addr) onlyAdmin public returns(bool){
        LockContractAddr = addr;
        return true;
    }
    
    function getLockContract() public view returns(address){
        return LockContractAddr;
    }
    
    function stake(address nodeAddr) public{
        //只能是质押合约来更新
        require(msg.sender == LockContractAddr);
        require(LockNodesIndexMap[nodeAddr].coinbase == address(0));
        require(BoeNodesIndexMap[nodeAddr].coinbase != address(0));
        LockNodes.push(nodeAddr);
        LockNodesIndexMap[nodeAddr].coinbase = nodeAddr;
        LockNodesIndexMap[nodeAddr].index = LockNodes.length-1;
    }
    
    function withdraw(address nodeAddr) public{
        //只能是质押合约来更新
        require(msg.sender == LockContractAddr);
        require(LockNodesIndexMap[nodeAddr].coinbase != address(0));
        uint256 index = LockNodesIndexMap[nodeAddr].index;
        LockNodes[index] = LockNodes[LockNodes.length-1];
        delete LockNodes[LockNodes.length-1];
        LockNodesIndexMap[LockNodes[index]].index = index;
        delete LockNodesIndexMap[nodeAddr];
        LockNodes.length--;
    }
    
    mapping (address => address) boetoholder;//boe节点地址=>持币地址
    mapping (address => address) holdertoboe;//持币地址=》boe地址
    //持币地址的设置需要双向确认
    function setHolderAddr(address holderAddr) public { //Boe地址调用此函数设置持币地址
        require(isBoeNode(msg.sender)); //只有boe地址能调用此函数
        require(holdertoboe[holderAddr] == address(0));//需要holderAddr未使用
        boetoholder[msg.sender]=holderAddr;
		emit SetHolderAddr(msg.sender,holderAddr);
    }
    
    //允许持币地址重复设置，重复设置时，取消前次的设置，
    //避免因误操作先设置了其他人的boe地址之后无法设置自己的boe地址
    function setHolderBoeAddr(address boeaddr) public{ //持币地址调用此函数，设置要支持的boe地址
        require(isBoeNode(boeaddr));
        //require(holdertoboe[msg.sender] == address(0));//取消此限制，避免误操作后无法设置的问题
        require(boetoholder[boeaddr] == msg.sender);
        holdertoboe[msg.sender] = boeaddr;
    }

    function fetchAllHolderAddrs() public view returns(address[] memory,
        address[] memory){
            address[] memory nodes = getAllBoesAddrs();
            address[] memory holders = new address[](nodes.length);
            for(uint i = 0; i < nodes.length; i++){
                holders[i] = getHolderAddr(nodes[i]);
            }
            return (nodes,holders);
    }

    function getHolderAddr(
        address  boeaddr
    )  public view returns (
        address 
    ){
        if (holdertoboe[boetoholder[boeaddr]] == boeaddr){
            return boetoholder[boeaddr];
        }
        return boeaddr;
    }
    
    event SetHolderAddr(
        address indexed coinBase,
        address indexed holderAddr
    );
}