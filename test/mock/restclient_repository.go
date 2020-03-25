package mock

import (
	"github.com/juanimoli/piccadilly/pkg/domain/model"
	"github.com/stretchr/testify/mock"
)

type RestClientRepositoryMock struct {
	mock.Mock
}

func (r *RestClientRepositoryMock) GetUsers(userGroupID string) ([]model.User, error) {
	args := r.Called(userGroupID)
	return args.Get(0).([]model.User), args.Error(1)
}