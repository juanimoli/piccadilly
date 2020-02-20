package random

import (
	"bytes"
	"encoding/json"
	"github.com/juanimoli/piccadilly/api/controller"
	"github.com/juanimoli/piccadilly/api/http"
	"github.com/juanimoli/piccadilly/api/model"
	"io/ioutil"
	"math/rand"
	net "net/http"
	"os"
	"strings"
)

func CreatePostController() controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/slack/random",
		Body:   CreatePostBody(),
	}
}

func CreatePostBody() http.Handler {
	return func(ctx *http.Context) {
		var reviewRequest model.ReviewRequest
		err := ctx.ReadBody(&reviewRequest)

		if err != nil {
			ctx.AbortTransactionWithError(http.CreateInternalError())
		}

		params := strings.Fields(reviewRequest.Text)
		if len(params) != 2 {
			ctx.AbortTransactionWithError(http.CreateBadRequestError("wrong number of parameters"))
		}

		userGroupId := params[0]

		resp, err := net.Get("https://slack.com/api/usergroups.users.list?token=" + os.Getenv("SECRET_SLACK_TOKEN") + "&usergroup=" + userGroupId + "&pretty=1")

		if err != nil {
			ctx.AbortTransactionWithError(http.CreateInternalError())
		}

		defer resp.Body.Close()

		if resp.StatusCode == net.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
			}

			var slackUserGroupResponse model.SlackUserGroup
			err = json.Unmarshal(bodyBytes, &slackUserGroupResponse)

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
			}

			if !slackUserGroupResponse.Ok {
				ctx.AbortTransactionWithError(http.CreateInternalError())
			}

			selected := slackUserGroupResponse.Users[rand.Intn(len(slackUserGroupResponse.Users))]

			bodyBytes, err = json.Marshal(model.SlackMessage{Message: selected + " have been chosen for review"})

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
			}

			net.Post(reviewRequest.ResponseUrl, "application/json", bytes.NewReader(bodyBytes))
		}
	}
}
