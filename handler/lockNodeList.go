package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/model"
	"github.com/hpb-project/votedapp-server/task"
	"net/http"
	"strings"
)

func GetLockedList(c *gin.Context) {
	t := task.BoeListTask()
	if t == nil {
		ResponseError(c, http.StatusInternalServerError, fmt.Sprintf("server not ready"))
		return
	}
	allNodes := t.GetNodes()
	allName := t.GetName()

	res := make([]*model.LockNodeInfo, 0)
	for _, info := range allNodes {
		r := &model.LockNodeInfo{
			Name:       fmt.Sprintf("node_%s", info.Coinbase[:10]),
			Coinbase:   info.Coinbase,
			LockNumber: info.LockNumber,
			LockAddr:   info.LockAddr,
		}
		name, exist := allName[strings.ToLower(r.Coinbase)]
		if exist {
			r.Name = name
		}
		res = append(res, r)
	}

	ResponseSuccess(c, res)
}
