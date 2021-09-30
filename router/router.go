package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hpb-project/votedapp-server/handler"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.POST("/nodename:addr", handler.GetNodeName)
}
