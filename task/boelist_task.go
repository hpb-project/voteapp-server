package task

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hpb-project/votedapp-server/config"
	"github.com/hpb-project/votedapp-server/contracts"
	"github.com/hpb-project/votedapp-server/db"
	log "github.com/sirupsen/logrus"
	"math/big"
	"sync"
	"time"
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
}

func newBoeListTask() (*BoeListRefresh, error) {
	var err error
	b := &BoeListRefresh{}
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

func (b *BoeListRefresh) loop() error {
	defer b.wg.Done()

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
				ticker.Reset(time.Second * 5)
			}
		case <-ticker.C:
			log.Info("boelist get new info")
			go b.getNewInfo()
		}
	}
}

func (b *BoeListRefresh) getNewInfo() {
	if b.working {
		return
	}
	b.working = true
	defer func() { b.working = false }()

	proxyAddr := common.HexToAddress(b.conf.ProxyAddr)
	proxy, err := contracts.NewProxy(proxyAddr, b.client)
	if err != nil {
		log.Errorf("get contract proxy failed, err:%s\n", err)
		return
	}
	defaultOpt := &bind.CallOpts{}
	nodesAddr, lockAddr, _, err := proxy.Getcontract(defaultOpt)
	if err != nil {
		log.Errorf("get contract address from proxy failed, err:%s\n", err)
		return
	}
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

	boelist, err := nodeContract.GetAllBoesAddrs(defaultOpt)
	if err != nil {
		log.Errorf("get allBoe failed, err:%s\n", err)
		return
	}

	for _, node := range boelist {
		n := &info{}
		n.nodeAddr = node
		nodeinfo[node] = n
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
	nodedb.Drop()
	nodedb.NewTable()

	records := make([]*db.BoeNode, 0)

	for _, value := range nodeinfo {
		r := &db.BoeNode{
			Coinbase:   value.nodeAddr.String(),
			LockAddr:   value.lockedAddr.String(),
			LockNumber: value.lockedAmount,
			HolderAddr: value.holderAddr.String(),
			VoteNum:    value.voteAmount,
		}
		records = append(records, r)
	}
	err = nodedb.BatchCreate(records)
	if err != nil {
		log.Errorf("node db batch create failed, err:%s\n", err)
	}
	log.Info("update node info finished")
}
