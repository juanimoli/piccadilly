package model

type User struct {
	ID string `json:"user_id",validate:"required"`

	Name string `json:"user_name",validate:"required"`
}

type ReviewRequest struct {
	ChannelId string `json:"channel_id",validate:"required"`

	*User

	Text string `json:"text",validate:"required"`

	ResponseUrl string `json:"response_url",validate:"required"`
}
