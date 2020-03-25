package random

import (
	"regexp"
	"strings"

	"github.com/juanimoli/piccadilly/pkg/domain/http"
	"github.com/juanimoli/piccadilly/pkg/domain/model"
)

func Map(ctx http.Reader) (model.ReviewRequest, error) {
	//map from response
	var existsChanId, existsText, existsRespUrl, existsUserName, existsUserId bool
	reviewRequest := model.ReviewRequest{}
	reviewRequest.ChannelId, existsChanId = ctx.GetFormData("channel_id")
	text, existsText := ctx.GetFormData("text")
	reviewRequest.ResponseUrl, existsRespUrl = ctx.GetFormData("response_url")
	reviewRequest.User = &model.User{}
	reviewRequest.User.Name, existsUserName = ctx.GetFormData("user_name")
	reviewRequest.User.ID, existsUserId = ctx.GetFormData("user_id")

	//check response
	if !existsChanId || !existsText || !existsRespUrl || !existsUserName || !existsUserId {
		//TODO wrap error
		return reviewRequest, http.CreateBadRequestError("missing parameters")
	}

	//check params
	params := strings.Fields(text)
	if len(params) != 2 {
		//TODO wrap error
		return reviewRequest, http.CreateBadRequestError("wrong number of parameters")
	}

	reviewRequest.PullRequestUrl = params[1]

	//map get params
	//TODO check regex
	r := regexp.MustCompile(`^<!subteam\^(.*)\|.*>$`).FindStringSubmatch(params[0])
	if len(r) != 2 {
		return reviewRequest, http.CreateBadRequestError("user group malformed")
	}
	reviewRequest.UserGroupId = r[1]

	return reviewRequest, nil
}
