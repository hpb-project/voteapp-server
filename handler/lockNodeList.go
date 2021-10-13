package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/db"
	"github.com/hpb-project/votedapp-server/model"
	"net/http"
	"strings"
)

func GetLockedList(c *gin.Context) {
	nodeTable := db.BoeNode{}
	nameTable := db.NodeName{}

	allNodes, err := nodeTable.GetAll()
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, fmt.Sprintf("got err:%s", err.Error()))
		return
	}

	res := make([]*model.LockNodeInfo, 0)
	for _, info := range allNodes {
		r := &model.LockNodeInfo{
			Name:       fmt.Sprintf("node_%s", info.Coinbase[:10]),
			Coinbase:   info.Coinbase,
			LockNumber: info.LockNumber,
			LockAddr:   info.LockAddr,
		}
		name, err := nameTable.GetNameByCoinbase(strings.ToLower(r.Coinbase))
		if err != nil {
			r.Name = name
		}
		res = append(res, r)
	}

	ResponseSuccess(c, res)
}
