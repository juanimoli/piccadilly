package usecase

import (
	"fmt"
	"github.com/juanimoli/piccadilly/pkg/domain/http"

	"github.com/juanimoli/piccadilly/pkg/domain/model"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
)

type GetRandomUserUseCase func(model.ReviewRequest) (model.SlackMessage, error)

type IntRagedRandom interface {
	RandomNumber(intRage int) int
}

func CreateGetRandomUserUseCase(clientRepository repository.RestClientRepository, random IntRagedRandom) GetRandomUserUseCase {
	return func(request model.ReviewRequest) (model.SlackMessage, error) {
		users, err := clientRepository.GetUsers(request.UserGroupId)

		if err != nil {
			return model.SlackMessage{}, err
		}

		myself := -1
		for i, user := range users {
			if user.ID == request.ID {
				myself = i
			}
		}

		if myself != -1 {
			users[len(users)-1], users[myself] = users[myself], users[len(users)-1]
			users = users[:len(users)-1]
		}

		if len(users) == 0 {
			return model.SlackMessage{}, http.CreateBadRequestError("no users available")
		}
		selected := users[random.RandomNumber(len(users))]
		return model.SlackMessage{
			Message:      fmt.Sprintf("<@%s> has been selected to review %s pull request", selected.ID, request.PullRequestUrl),
			Channel:      request.ChannelId,
			ResponseType: "in_channel",
		}, nil
	}
}
