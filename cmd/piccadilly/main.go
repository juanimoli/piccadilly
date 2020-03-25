package main

import (
	"fmt"
	random2 "github.com/juanimoli/piccadilly/pkg/infra/random"
	"math/rand"
	"time"

	"github.com/juanimoli/piccadilly/pkg/data/usecase"
	"github.com/juanimoli/piccadilly/pkg/infra/gin/engine"
	"github.com/juanimoli/piccadilly/pkg/infra/heimdall"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
	"github.com/juanimoli/piccadilly/pkg/presenter/health"
	"github.com/juanimoli/piccadilly/pkg/presenter/random"
)

func main() {
	e := engine.New()

	randomInstance := random2.CreateIntRangedRandom(rand.New(rand.NewSource(time.Now().UnixNano())))
	heimdallClient := heimdall.CreateGetClient()
	slackRestClientRepository := repository.CreateSlackRestClientRepository(heimdallClient)
	getRandomUseCase := usecase.CreateGetRandomUserUseCase(slackRestClientRepository, randomInstance)

	//********** Health handlers ***********//
	e.GET("/ping", health.CreateGetBody())

	//********** Slack handlers ***********//
	e.POST("/slack/random", random.CreatePostBody(getRandomUseCase))

	if err := e.Run("3000"); err != nil {
		fmt.Println(err.Error())
	}
}
