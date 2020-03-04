package heimdall

import (
	"github.com/gojektech/heimdall/httpclient"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
	"net/http"
	"time"
)

func CreateGetClient() repository.HttpGet {
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(1000 * time.Millisecond))
	return func(url string) (response *http.Response, err error) {
		response, err = client.Get(url, nil)
		return
	}
}
