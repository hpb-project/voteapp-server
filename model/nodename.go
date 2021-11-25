package model

type NodeNameQueryParam struct {
	Coinbase string `json:"coinbase"`
}

type NodeName struct {
	Name    string `json:"name"`
	NameEng string `json:"name_eng"`
}

type VoteNodeInfo struct {
	Name       string `json:"name"`       // name
	NameEng    string `json:"name_eng"`   // english name
	Coinbase   string `json:"coinbase"`   // addr
	VoteNumber uint64 `json:"voteNumber"` // voteNum
}

type LockNodeInfo struct {
	Name       string `json:"name"`       // name
	NameEng    string `json:"name_eng"`   // english name
	Coinbase   string `json:"coinbase"`   // coinbase
	LockAddr   string `json:"lockAddr"`   // lockAddr
	LockNumber uint64 `json:"lockNumber"` // lockNum
}
