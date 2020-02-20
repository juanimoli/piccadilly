package model

type SlackMessage struct {
	Message string `json:"text"`

	Channel string `json:"channel"`

	ReplyBroadcast bool `json:"reply_broadcast"`

	ThreadTs string `json:"thread_ts"`

	DeleteOriginal string `json:"delete_original"`
}
