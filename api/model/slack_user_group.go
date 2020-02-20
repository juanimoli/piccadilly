package model

type SlackUserGroup struct {
	Ok bool `validate:"required"`

	Users []string `json:"user",validate:"required"`
}
