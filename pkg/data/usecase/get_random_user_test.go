package usecase_test

import (
	"errors"
	"github.com/juanimoli/piccadilly/pkg/data/usecase"
	"github.com/juanimoli/piccadilly/pkg/domain/model"
	"github.com/juanimoli/piccadilly/test/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRandomUserGivenAClientRepositoryThatReturnsError_WhenExecutingUseCase_ThenErrorIsReturned(t *testing.T) {
	userGroupId := "test"
	mockRepository := new(mock.RestClientRepositoryMock)
	mockRepository.On("GetUsers", userGroupId).Return([]model.User{}, errors.New("test"))
	getRandomUserUseCase := usecase.CreateGetRandomUserUseCase(mockRepository, nil)

	_, err := getRandomUserUseCase(model.ReviewRequest{
		UserGroupId: userGroupId,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "test", err.Error())
}

func TestGetRandomUserGivenAClientRepositoryThatReturnsEmptyList_WhenExecutingUseCase_ThenErrorIsReturned(t *testing.T) {
	userGroupId := "test"
	mockRepository := new(mock.RestClientRepositoryMock)
	mockRepository.On("GetUsers", userGroupId).Return([]model.User{}, nil)
	getRandomUserUseCase := usecase.CreateGetRandomUserUseCase(mockRepository, nil)

	_, err := getRandomUserUseCase(model.ReviewRequest{
		UserGroupId: userGroupId,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no users available", err.Error())
}

func TestGetRandomUserGivenAClientRepositoryThatReturnsOnlyMyself_WhenExecutingUseCase_ThenErrorIsReturned(t *testing.T) {
	userGroupId := "test"
	userId := "test"
	anotherUserId := "test2"
	mockRepository := new(mock.RestClientRepositoryMock)
	mockRepository.On("GetUsers", userGroupId).Return([]model.User{
		{
			ID: userId,
		},
		{
			ID: anotherUserId,
		},
	}, nil)
	mockRandom := new(mock.IntRangedRandomMock)
	mockRandom.On("RandomNumber", 1).Return(0)
	getRandomUserUseCase := usecase.CreateGetRandomUserUseCase(mockRepository, mockRandom)

	result, err := getRandomUserUseCase(model.ReviewRequest{
		PullRequestUrl: "test",
		UserGroupId:    userGroupId,
		User: &model.User{
			ID: userId,
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, "<@test2> has been selected to review test pull request", result.Message)
}
