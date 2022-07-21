package task

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/votedapp-server/config"
	"github.com/hpb-project/votedapp-server/contracts"
	"github.com/hpb-project/votedapp-server/db"
	"github.com/hpb-project/votedapp-server/model"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
	"sync"
	"time"
)

const (
	hpbNodeContract   = "0x451d785A0379E637d17C1C0E96cA150168A5Ab9A"
	hpbLockContract   = "0x74054455F954E1DDAf0694d906e4e68D63a33A18"
	hpbVoteContract   = "0x35a3445C0ca0B01B7CEA2F867D762f6410c9e952"
	hpbProxyContract  = "0x60a3698bE1493da2065E6F84B2E77B5b5D201D5D"
	hpbFilterContract = "0x303632d9E171a130B85ee1a23cD5902e5a2F7A51"
)

// used to refresh boelist
type BoeListRefresh struct {
	conf    *config.Config
	client  *ethclient.Client
	quit    chan struct{}
	stop    chan struct{}
	start   chan struct{}
	working bool
	wg      sync.WaitGroup

	cacheName map[string]model.NodeName
	cache     []*db.BoeNode
	cacheLock sync.RWMutex
}

var (
	boelistTask *BoeListRefresh
)

func BoeListTask() *BoeListRefresh {
	return boelistTask
}

func newBoeListTask() (*BoeListRefresh, error) {
	if boelistTask != nil {
		return boelistTask, nil
	}
	var err error
	b := &BoeListRefresh{cacheName: make(map[string]model.NodeName), cache: make([]*db.BoeNode, 0)}
	b.conf = config.GetConfig()
	b.client, err = ethclient.Dial(b.conf.RPC)
	if err != nil {
		return nil, err
	}
	b.quit = make(chan struct{}, 1)
	b.stop = make(chan struct{}, 1)
	b.start = make(chan struct{}, 1)
	b.wg.Add(1)
	go func() { err = b.loop() }()

	if err != nil {
		return nil, err
	}

	boelistTask = b
	return b, nil
}

func (b *BoeListRefresh) Start() error {
	b.start <- struct{}{}
	return nil
}

func (b *BoeListRefresh) Stop() error {
	b.stop <- struct{}{}
	return nil
}

func (b *BoeListRefresh) GetNodes() []db.BoeNode {
	b.cacheLock.RLock()
	defer b.cacheLock.RUnlock()
	nodes := make([]db.BoeNode, 0)
	for _, d := range b.cache {
		nodes = append(nodes, *d)
	}
	return nodes
}

func (b *BoeListRefresh) GetName() map[string]model.NodeName {
	b.cacheLock.RLock()
	defer b.cacheLock.RUnlock()
	namemap := make(map[string]model.NodeName)
	for k, v := range b.cacheName {
		namemap[k] = v
	}

	return namemap
}

func (b *BoeListRefresh) loop() error {
	defer b.wg.Done()
	loopCycle := time.Second * 15
	ticker := time.NewTicker(time.Second * 999999)
	defer ticker.Stop()
	for {
		log.Debug("boelist loop wait")
		select {
		case _, ok := <-b.stop:
			if ok {
				ticker.Reset(time.Hour * 999999)
			}
		case _, ok := <-b.start:
			log.Info("boelist task started")
			if ok {
				log.Debug("boelist loop reset ticket to 5s")
				ticker.Reset(loopCycle)
				go b.getNewInfo()
				go b.updateName()
			}
		case <-ticker.C:
			log.Info("boelist get new info")
			go b.getNewInfo()
			go b.updateName()
		}
	}
}

//
//func (b *BoeListRefresh) monitor() {
//	proxyAddr := common.HexToAddress(b.conf.ProxyAddr)
//	proxy, err := contracts.NewProxy(proxyAddr, b.client)
//	if err != nil {
//		log.Errorf("get contract proxy failed, err:%s\n", err)
//		return
//	}
//
//	defaultOpt := &bind.CallOpts{}
//
//	nodesAddr, lockAddr, voteAddr, err := proxy.Getcontract(defaultOpt)
//	if err != nil {
//		log.Errorf("get contract address from proxy failed, err:%s\n", err)
//		return
//	}
//	nodeContract, err := contracts.NewNodes(nodesAddr, b.client)
//	if err != nil {
//		log.Errorf("get contract nodes failed, err:%s\n", err)
//		return
//	}
//	loclContract, err := contracts.NewLock(lockAddr, b.client)
//	if err != nil {
//		log.Errorf("get contract lock failed, err:%s\n", err)
//		return
//	}
//	voteContract, err := contracts.NewVote(voteAddr, b.client)
//	if err != nil {
//		log.Errorf("get contract vote failed, err:%s\n", err)
//		return
//	}
//	wopt := &bind.WatchOpts{}
//	newBoe := make(chan *contracts.NodesAddBoeNode)
//	addNodeSub, err := nodeContract.WatchAddBoeNode(wopt, newBoe, []common.Address{}, [][32]byte{})
//	if err != nil {
//		log.Errorf("watch addBoeNode failed, err:%s\n", err)
//	}
//
//	delBoe := make(chan *contracts.NodesDeleteBoeNode)
//	delNodeSub, err := nodeContract.WatchDeleteBoeNode(wopt, delBoe, []common.Address)
//	if err != nil {
//		log.Errorf("watch delBoeNode failed, err:%s\n", err)
//	}
//	doVote := make(chan *contracts.VoteDoVoted)
//	voteSub, err := voteContract.WatchDoVoted(wopt, doVote, []*big.Int{}, []common.Address, []common.Address)
//	if err != nil {
//		log.Errorf("watch doVote failed, err:%s\n", err)
//	}
//
//	for {
//		select {
//		case <-addNodeSub.Err():
//
//		}
//	}
//
//}

func (b *BoeListRefresh) updateName() {
	nametable := &db.NodeName{}
	nameInfo, err := nametable.GetAllInfo()
	if err != nil {
		log.Errorf("can't get node name info, err:%s\n", err)
		return
	}
	b.cacheLock.Lock()
	b.cacheName = make(map[string]model.NodeName)
	for _, info := range nameInfo {
		b.cacheName[strings.ToLower(info.Coinbase)] = model.NodeName{NameEng: info.NameEng, Name: info.Name}
	}
	b.cacheLock.Unlock()
}

func (b *BoeListRefresh) getNewInfo() {
	if b.working {
		return
	}
	b.working = true
	defer func() { b.working = false }()

	proxyAddr := common.HexToAddress(hpbProxyContract)
	proxy, err := contracts.NewProxy(proxyAddr, b.client)
	if err != nil {
		log.Errorf("get contract proxy failed, err:%s\n", err)
		return
	}

	defaultOpt := &bind.CallOpts{}
	if block, err := b.client.BlockNumber(context.Background()); err != nil {
		defaultOpt.BlockNumber = big.NewInt(int64(block))
	}

	nodesAddr, lockAddr := common.HexToAddress(hpbNodeContract), common.HexToAddress(hpbLockContract)

	nodeContract, err := contracts.NewNodes(nodesAddr, b.client)
	if err != nil {
		log.Errorf("get contract nodes failed, err:%s\n", err)
		return
	}
	type info struct {
		nodeAddr     common.Address // 节点地址
		holderAddr   common.Address // 持币地址
		lockedAddr   common.Address // 锁仓账户
		lockedAmount *big.Int       // 锁仓金额
		voteAmount   *big.Int       // 投票数量
	}
	var nodeinfo = make(map[common.Address]*info)
	var allnodelist = make([]*info, 0)

	boelist, err := nodeContract.GetAllBoesAddrs(defaultOpt)
	if err != nil {
		log.Errorf("get allBoe failed, err:%s\n", err)
		return
	}

	for _, node := range boelist {
		n := &info{}
		n.nodeAddr = node
		nodeinfo[node] = n
		allnodelist = append(allnodelist, n)
	}

	lockContract, err := contracts.NewLock(lockAddr, b.client)
	if err != nil {
		log.Errorf("get contract lock failed, err:%s\n", err)
		return
	}

	{ // get lock info.
		lockedlist, err := nodeContract.GetAlllockNode(defaultOpt)
		if err != nil {
			log.Errorf("get allLocked failed, err:%s\n", err)
			return
		}

		for _, addr := range lockedlist {

			if lockedAddr, lockedAmount, err := lockContract.FetchLockInfoByNodeAddr(defaultOpt, addr); err != nil {
				log.Errorf("fetch locked info failed, err:%s\n", err)
				return
			} else {
				n, exist := nodeinfo[addr]
				if !exist {
					continue
				}
				n.lockedAddr = lockedAddr
				n.lockedAmount = lockedAmount
			}
		}

	}

	{ // get vote info.
		addrlist, amountlist, err := proxy.FetchAllVoteResult(defaultOpt)
		if err != nil {
			log.Errorf("fetch all vote result failed, err:%s\n", err)
			return
		}
		for i, addr := range addrlist {
			if n, exist := nodeinfo[addr]; !exist {
				continue
			} else {
				n.voteAmount = amountlist[i]
			}
		}
	}

	{ // get holder address.
		nodelist, holderlist, err := proxy.FetchAllHolderAddrs(defaultOpt)
		if err != nil {
			log.Errorf("fetch all holder failed, err:%s\n", err)
			return
		}
		for i, addr := range nodelist {
			if n, exist := nodeinfo[addr]; !exist {
				continue
			} else {
				n.holderAddr = holderlist[i]
			}
		}
	}

	nodedb := &db.BoeNode{}
	nodedb.NewTable()

	ether := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	records := make([]*db.BoeNode, 0)
	for _, value := range allnodelist {
		r := &db.BoeNode{
			Coinbase:   value.nodeAddr.String(),
			LockAddr:   value.lockedAddr.String(),
			LockNumber: 0,
			HolderAddr: value.holderAddr.String(),
			VoteNum:    0,
		}
		if value.lockedAmount != nil {
			r.LockNumber = new(big.Int).Quo(value.lockedAmount, ether).Uint64()
		}
		if value.voteAmount != nil {
			r.VoteNum = new(big.Int).Quo(value.voteAmount, ether).Uint64()
		}
		records = append(records, r)
	}

	b.cacheLock.Lock()
	b.cache = records
	b.cacheLock.Unlock()

	err = nodedb.RefreshAll(records)
	if err != nil {
		log.Errorf("node db batch create failed, err:%s\n", err)
	}
	log.Info("update node info finished")
}
