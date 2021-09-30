package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/db"
	"net/http"
	"strings"
)

func GetNodeName(c *gin.Context) {
	addr := c.Param("addr")
	if len(addr) != 42 {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("invalid param"))
		return
	}
	d := &db.NodeName{}
	name, err := d.GetNameByCoinbase(strings.ToLower(addr))
	if err != nil {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("got err:%s", err.Error()))
	} else {
		ResponseSuccess(c, name)
	}
}
