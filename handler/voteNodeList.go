package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/model"
	"github.com/hpb-project/votedapp-server/task"
	"net/http"
	"strings"
)

func GetVoteList(c *gin.Context) {
	t := task.BoeListTask()
	if t == nil {
		ResponseError(c, http.StatusInternalServerError, fmt.Sprintf("server not ready"))
		return
	}
	allNodes := t.GetNodes()
	allName := t.GetName()

	res := make([]*model.VoteNodeInfo, 0)
	for _, info := range allNodes {
		if info.LockNumber <= 0 {
			continue
		}
		r := &model.VoteNodeInfo{
			Name:       fmt.Sprintf("node_%s", info.Coinbase[:10]),
			Coinbase:   info.Coinbase,
			VoteNumber: info.VoteNum,
		}
		name, exist := allName[strings.ToLower(r.Coinbase)]
		if exist {
			r.Name = name
		}
		res = append(res, r)
	}

	ResponseSuccess(c, res)
}
