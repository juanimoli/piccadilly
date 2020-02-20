package random

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/juanimoli/piccadilly/api/controller"
	"github.com/juanimoli/piccadilly/api/http"
	"github.com/juanimoli/piccadilly/api/model"
	"io/ioutil"
	"log"
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
		var existsChanId, existsText, existsRespUrl, existsUserName, existsUserId bool
		reviewRequest := model.ReviewRequest{}
		reviewRequest.ChannelId, existsChanId = ctx.GetFormData("channel_id")
		reviewRequest.Text, existsText = ctx.GetFormData("text")
		reviewRequest.ResponseUrl, existsRespUrl = ctx.GetFormData("response_url")
		reviewRequest.User = &model.User{}
		reviewRequest.User.Name, existsUserName = ctx.GetFormData("user_name")
		reviewRequest.User.ID,existsUserId = ctx.GetFormData("user_id")

		if !existsChanId || !existsText || !existsRespUrl || !existsUserName || !existsUserId {
			ctx.AbortTransactionWithError(http.CreateInternalError())
			//TODO not 400 capo, mandale un mensaje que se equivoco de params
			log.Fatal("Missing parameters")
			return
		}

		fmt.Println("1")

		params := strings.Fields(reviewRequest.Text)
		if len(params) != 2 {
			//TODO not 400 capo, mandale un mensaje que se equivoco de params
			ctx.AbortTransactionWithError(http.CreateBadRequestError("wrong number of parameters"))
			log.Fatal("Wrong number of parameters")
			return
		}

		fmt.Println("2")

		userGroupId := params[0]

		resp, err := net.Get("https://slack.com/api/usergroups.users.list?token=" + os.Getenv("SECRET_SLACK_TOKEN") + "&usergroup=" + userGroupId + "&pretty=1")

		if err != nil {
			ctx.AbortTransactionWithError(http.CreateInternalError())
			log.Fatal(err)
			return
		}

		fmt.Println("3")

		defer resp.Body.Close()

		if resp.StatusCode == net.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)

			fmt.Println("4")

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
				log.Fatal(err)
				return
			}

			fmt.Println("4")

			var slackUserGroupResponse model.SlackUserGroup
			err = json.Unmarshal(bodyBytes, &slackUserGroupResponse)

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
				log.Fatal(err)
				return
			}

			fmt.Println("5")

			if !slackUserGroupResponse.Ok {
				ctx.AbortTransactionWithError(http.CreateInternalError())
				log.Fatal(err)
				return
			}

			selected := slackUserGroupResponse.Users[rand.Intn(len(slackUserGroupResponse.Users))]

			bodyBytes, err = json.Marshal(model.SlackMessage{Message: selected + " have been chosen for review"})

			if err != nil {
				ctx.AbortTransactionWithError(http.CreateInternalError())
				log.Fatal(err)
				return
			}

			fmt.Println("6")

			net.Post(reviewRequest.ResponseUrl, "application/json", bytes.NewReader(bodyBytes))
		}
	}
}
