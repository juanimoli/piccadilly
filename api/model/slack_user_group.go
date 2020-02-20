package model

type SlackUserGroup struct {
	Ok bool `validate:"required"`

	Users []string `json:"users",validate:"required"`
}
