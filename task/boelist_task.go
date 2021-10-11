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
}

func newBoeListTask() (*BoeListRefresh, error) {
	var err error
	b := &BoeListRefresh{}
	b.conf = config.GetConfig()
	b.client, err = ethclient.Dial(b.conf.RPC)
	if err != nil {
		return nil, err
	}
	b.quit = make(chan struct{})
	b.stop = make(chan struct{})
	b.start = make(chan struct{})
	return &BoeListRefresh{}, nil
}

func (b *BoeListRefresh) Start() error {

	return nil
}

func (b *BoeListRefresh) Stop() error {
	return nil
}

func (b *BoeListRefresh) loop() error {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case _, closed := <-b.stop:
			if !closed {
				ticker.Reset(time.Hour * 999999)
			}
		case _, closed := <-b.start:
			if !closed {
				ticker.Reset(time.Second * 5)
			}
		case <-ticker.C:
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
	for _, value := range nodeinfo {
		r := &db.BoeNode{
			Coinbase:   value.nodeAddr.String(),
			LockAddr:   value.lockedAddr.String(),
			LockNumber: value.lockedAmount,
			HolderAddr: value.holderAddr.String(),
			VoteNum:    value.voteAmount,
		}
		r.Create()
	}
}
