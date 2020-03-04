package usecase

import (
	"fmt"
	"github.com/juanimoli/piccadilly/pkg/domain/model"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
	"math/rand"
)

type GetRandomUserUseCase func(model.ReviewRequest) (model.SlackMessage, error)

func CreateGetRandomUserUseCase(clientRepository repository.RestClientRepository) GetRandomUserUseCase {
	return func(request model.ReviewRequest) (model.SlackMessage, error) {
		users, err := clientRepository.GetUsers(request.UserGroupId)

		if err != nil {
			return model.SlackMessage{}, err
		}

		//TODO check que no salga siempre el mismo y que no salgas vos mismo
		selected := users[rand.Intn(len(users))]
		return model.SlackMessage{
			Message:      fmt.Sprintf("<@%s> has been selected to review %s pull request", selected, request.PullRequestUrl),
			Channel:      request.ChannelId,
			ResponseType: "in_channel",
			//TODO CHEQUEAME ESTO BRE
			DeleteOriginal: "true",
		}, nil
	}
}
