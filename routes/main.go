package routes

import (
	"github.com/gin-gonic/gin"
)

type IGinGonicRouter interface {
	InitRouter(g *gin.Engine)
}

type Router struct{}

func (router *Router) InitRouter(g *gin.Engine) {
	versionGroup := g.Group("/v1")
	InitCurrencyRoute(versionGroup)
}

func GinGonicRouter() IGinGonicRouter {
	return &Router{}
}
