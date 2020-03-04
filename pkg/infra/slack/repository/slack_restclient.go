package repository

import (
	"encoding/json"
	"github.com/juanimoli/piccadilly/pkg/domain/model"
	model2 "github.com/juanimoli/piccadilly/pkg/infra/slack/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type HttpGet func(url string) (*http.Response, error)

type RestClientRepository interface {
	GetUsers(userGroupID string) ([]model.User, error)
}

type slackRestClientRepository struct {
	HttpGet HttpGet
}

func CreateSlackRestClientRepository(get HttpGet) RestClientRepository {
	return &slackRestClientRepository{
		HttpGet: get,
	}
}

const (
	slackUserGroupURL               = "https://slack.com/api/usergroups.users.list"
	slackUserGroupURLTokenParam     = "token"
	slackUserGroupURLUserGroupParam = "userGroup"
)

func (s slackRestClientRepository) GetUsers(userGroupID string) ([]model.User, error) {
	u, err := url.Parse(slackUserGroupURL)

	if err != nil {
		return []model.User{}, err
	}

	q := u.Query()
	q.Set(slackUserGroupURLTokenParam, os.Getenv("SECRET_SLACK_TOKEN"))
	q.Set(slackUserGroupURLUserGroupParam, userGroupID)
	u.RawQuery = q.Encode()

	res, err := s.HttpGet(u.String())
	if err != nil {
		return []model.User{}, err
	}

	//check response
	if res.StatusCode == http.StatusOK {
		//read response
		bodyBytes, err := ioutil.ReadAll(res.Body)

		//check read response
		if err != nil {
			return []model.User{}, err
		}

		//map response to model
		var slackUserGroupResponse model2.SlackUserGroup
		err = json.Unmarshal(bodyBytes, &slackUserGroupResponse)

		//check map
		if err != nil || !slackUserGroupResponse.Ok {
			return []model.User{}, err
		}

		return mapSlackUserGroup(slackUserGroupResponse), nil
	} else {
		return []model.User{}, err
	}
}

func mapSlackUserGroup(slackUserGroup model2.SlackUserGroup) []model.User {
	var users []model.User
	for _, user := range slackUserGroup.Users {
		u := model.User{
			ID: user,
		}
		users = append(users, u)
	}
	return users
}
