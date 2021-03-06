package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/handler"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/nodename/:addr", handler.GetNodeName)
	v1.GET("/nodenameeng/:addr", handler.GetNodeEngName)
	v1.GET("/locklist/", handler.GetLockedList)
	v1.GET("/votelist/", handler.GetVoteList)
}
