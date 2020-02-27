package model

type SlackMessage struct {
	Message string `json:"text"`

	Channel string `json:"channel"`

	ResponseType string `json:"response_type"`

	DeleteOriginal string `json:"delete_original"`
}
