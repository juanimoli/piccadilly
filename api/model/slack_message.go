package model

type SlackMessage struct {
	Message string `json:"text"`

	ReplyBroadcast string `json:"reply_broadcast"`

	ThreadTs string `json:"thread_ts"`

	DeleteOriginal string `json:"delete_original"`
}
