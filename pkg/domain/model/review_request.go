package model

type User struct {
	ID string

	Name string
}

type ReviewRequest struct {
	ChannelId string

	*User

	UserGroupId string

	PullRequestUrl string

	ResponseUrl string
}
