package main

import (
	"fmt"
	"github.com/juanimoli/piccadilly/cmd/piccadilly/infra/engine/gin"
	"github.com/juanimoli/piccadilly/pkg/presenter/health"
	"github.com/juanimoli/piccadilly/pkg/presenter/random"
)

func main() {
	e := gin.New()

	e.GET("/ping", health.CreateGetBody())
	e.POST("/slack/random", random.CreatePostBody())

	if err := e.Run("8080"); err != nil {
		fmt.Println(err.Error())
	}
}
