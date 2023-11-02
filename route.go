package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Route interface {
	Pattern() string
	Handler(*gin.Context)
	Register(*gin.Engine)
}

type HelloRoute struct{}

var _ Route = (*HelloRoute)(nil)

func (*HelloRoute) Pattern() string {
	return "/hello"
}

func (*HelloRoute) Handler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello",
	})
}

func (h *HelloRoute) Register(g *gin.Engine) {
	g.GET(h.Pattern(), h.Handler)
}

func NewHelloRoute() *HelloRoute {
	return &HelloRoute{}
}

func ProvideRoute(a any) any {
	return fx.Annotate(a, fx.As(new(Route)), fx.ResultTags(`group:"routes"`))
}
