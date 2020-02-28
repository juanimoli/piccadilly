package repository

import "github.com/juanimoli/piccadilly/pkg/domain/model"

type UserGroupRepository interface {
	GetUsers(userGroupID string) ([]model.User, error)
}
