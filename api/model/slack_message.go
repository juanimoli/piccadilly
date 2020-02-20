package model

type SlackMessage struct {
	Message string `json:"text"`

	ResponseType string `json:"response_type"`

	DeleteOriginal string `json:"delete_original"`
}
