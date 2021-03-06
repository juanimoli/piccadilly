package random_test

import (
	"encoding/json"
	"testing"

	"github.com/juanimoli/piccadilly/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

//func TestGivenAReaderMockWithFullJSON_WhenThePostBodyIsExecuted_ThenTheBodyIsConsumed(t *testing.T) {
//	var result *model.ReviewRequest
//	readerMock := new(mock2.Reader)
//	readerMock.On("ReadBody", mock.MatchedBy(func(obj *model.ReviewRequest) bool {
//		json.Unmarshal([]byte(`{
//			"channel_id": "DSEEXL90S",
//			"user_id": "USV6FBMTR",
//			"user_name": "jimolina",
//			"text": "link ID"
//		}`), obj)
//		result = obj
//		return true
//	})).Return(nil)
//
//	ctx := &http.Context{
//		Reader:     readerMock,
//		Writer:     nil,
//		Middleware: nil,
//	}
//
//	body := random.CreatePostBody()
//
//	body(ctx)
//
//	assert.Equal(t, "DSEEXL90S", result.ChannelId)
//	assert.Equal(t, "USV6FBMTR", result.User.ID)
//	assert.Equal(t, "jimolina", result.User.Name)
//	assert.Equal(t, "link ID", result.Text)
//
//	readerMock.AssertExpectations(t)
//}

func TestForradaPapu(t *testing.T) {
	var result model.ReviewRequest
	err := json.Unmarshal([]byte(`{
			"channel_id": "DSEEXL90S",
			"user_id": "USV6FBMTR",
			"user_name": "jimolina",
			"response_url": "lakjsdadfsjkldfas",
			"text": "tuvieja en tanga"
		}`), &result)

	assert.Nil(t, err)
}
