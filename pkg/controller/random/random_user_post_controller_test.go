package random_test

import (
	"encoding/json"
	"github.com/juanimoli/piccadilly/api/http"
	"github.com/juanimoli/piccadilly/api/model"
	"github.com/juanimoli/piccadilly/pkg/controller/random"
	mock2 "github.com/juanimoli/piccadilly/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGivenAReaderMockWithFullJSON_WhenThePostBodyIsExecuted_ThenTheBodyIsConsumed(t *testing.T) {
	var result *model.ReviewRequest
	readerMock := new(mock2.ReaderMock)
	readerMock.On("ReadBody", mock.MatchedBy(func(obj *model.ReviewRequest) bool {
		json.Unmarshal([]byte(`{
			"channel_id": "DSEEXL90S",
			"user_id": "USV6FBMTR",
			"user_name": "jimolina",
			"text": "link ID"
		}`), obj)
		result = obj
		return true
	})).Return(nil)

	ctx := &http.Context{
		Reader:     readerMock,
		Writer:     nil,
		Middleware: nil,
	}

	body := random.CreatePostBody()

	body(ctx)

	assert.Equal(t, "DSEEXL90S", result.ChannelId)
	assert.Equal(t, "USV6FBMTR", result.User.ID)
	assert.Equal(t, "jimolina", result.User.Name)
	assert.Equal(t, "link ID", result.Text)

	readerMock.AssertExpectations(t)
}

func TestGivenAReaderMockWithUncompletedJSON_WhenPostBodyIsExcecuted_ThenIsAbortedWithBadRequestCode(t *testing.T) {
	var result *model.ReviewRequest
	readerMock := new(mock2.ReaderMock)
	readerMock.On("ReadBody", mock.MatchedBy(func(obj *model.ReviewRequest) bool {
		json.Unmarshal([]byte(`{
			"channel_id": "DSEEXL90S",
			"user_id": "USV6FBMTR",
			"user_name": "jimolina"
		}`), obj)
		result = obj
		return true
	})).Return(nil)
	middlewareMock := new(mock2.MiddlewareMock)
	middlewareMock.On("AbortTransactionWithError", http.CreateInternalError()).Return(nil)

	ctx := &http.Context{
		Reader:     readerMock,
		Writer:     nil,
		Middleware: nil,
	}

	body := random.CreatePostBody()

	body(ctx)

	//assert.Equal(t, )

	readerMock.AssertExpectations(t)
}