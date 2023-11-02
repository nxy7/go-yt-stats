package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ProvideRoute(NewHelloRoute),
			fx.Annotate(StartServer, fx.ParamTags(`group:"routes"`)),
		),
		fx.Invoke(func(*gin.Engine) {}),
	).Run()
}
