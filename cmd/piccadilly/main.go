package main

import (
	"fmt"
	"github.com/juanimoli/piccadilly/pkg/data/usecase"
	"github.com/juanimoli/piccadilly/pkg/infra/gin/engine"
	"github.com/juanimoli/piccadilly/pkg/infra/heimdall"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
	"github.com/juanimoli/piccadilly/pkg/presenter/health"
	"github.com/juanimoli/piccadilly/pkg/presenter/random"
)

func main() {
	e := engine.New()
	heimdallClient := heimdall.CreateGetClient()
	slackRestClientRepository := repository.CreateSlackRestClientRepository(heimdallClient)
	getRandomUseCase := usecase.CreateGetRandomUserUseCase(slackRestClientRepository)

	e.GET("/ping", health.CreateGetBody())
	e.POST("/slack/random", random.CreatePostBody(getRandomUseCase))

	if err := e.Run("8080"); err != nil {
		fmt.Println(err.Error())
	}
}
