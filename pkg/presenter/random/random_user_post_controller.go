package random

import (
	"github.com/juanimoli/piccadilly/pkg/data/usecase"
	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

func CreatePostBody(useCase usecase.GetRandomUserUseCase) http.Handler {
	return func(ctx *http.Context) {
		reviewRequest, err := Map(ctx)

		if err != nil {
			ctx.AbortTransactionWithError(err)
		}

		slackMessage, err := useCase(reviewRequest)

		if err != nil {
			ctx.AbortTransactionWithError(err)
		}

		ctx.WriteJson(200, slackMessage)
	}
}
