package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/db"
	"net/http"
	"strings"
)

func GetNodeName(c *gin.Context) {
	addr := c.Query("addr")

	if len(addr) != 42 {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("invalid param"))
		return
	}
	d := &db.NodeName{}
	name, _, err := d.GetNameByCoinbase(strings.ToLower(addr))
	if err != nil {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("got err:%s", err.Error()))
	} else {
		ResponseSuccess(c, name)
	}
}

func GetNodeEngName(c *gin.Context) {
	addr := c.Query("addr")

	if len(addr) != 42 {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("invalid param"))
		return
	}
	d := &db.NodeName{}
	_, eng, err := d.GetNameByCoinbase(strings.ToLower(addr))
	if err != nil {
		ResponseError(c, http.StatusBadRequest, fmt.Sprintf("got err:%s", err.Error()))
	} else {
		ResponseSuccess(c, eng)
	}
}
