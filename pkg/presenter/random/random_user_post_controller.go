package random

import (
	"github.com/juanimoli/piccadilly/pkg/domain/http"
)

func CreatePostBody() http.Handler {
	return func(ctx *http.Context) {
		/*
			reviewRequest, err := Map(ctx)

			if err != nil {
				ctx.AbortTransactionWithError(err)
			}

			//api call to get users from usergroup
			resp, err :=

			//check response
			if resp != nil {
				//TODO: check lista vacia
				//TODO: sacarme a mi mismo de la lista
				//get random from user list
				selected := slackUserGroupResponse.Users[rand.Intn(len(slackUserGroupResponse.Users))]

				//shoot response
				ctx.WriteJson(200, model.SlackMessage{
					Message:      fmt.Sprintf("<@%s> has been selected to review %s pull request", selected, params[1]),
					Channel:      reviewRequest.ChannelId,
					ResponseType: "in_channel",
					//TODO CHEQUEAME ESTO BRE
					DeleteOriginal: "true",
				})
			} else {

			}
		*/
	}
}
