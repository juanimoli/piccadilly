package repository

import "github.com/juanimoli/piccadilly/pkg/domain/model"

type SlackRepository interface {
	GetUsersByUserGroup(userGroupID string) ([]model.User, error)
}
