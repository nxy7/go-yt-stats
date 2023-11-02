package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type App struct{}

func (a App) Run() {
	fmt.Println("elo")
}

func StartServer(routes []Route, lc fx.Lifecycle) *gin.Engine {
	g := gin.Default()

	log.Println("Routes: ", routes)

	for _, r := range routes {
		r.Register(g)
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Println("Server startup")
				go g.Run()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Stop server")
				return nil
			},
		},
	)

	return g
}
