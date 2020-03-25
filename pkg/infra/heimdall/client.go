package heimdall

import (
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
	"github.com/juanimoli/piccadilly/pkg/infra/slack/repository"
)

func CreateGetClient() repository.HttpGet {
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(1000 * time.Millisecond))
	return func(url string) (response *http.Response, err error) {
		response, err = client.Get(url, nil)
		return
	}
}
