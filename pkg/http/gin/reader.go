package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanimoli/piccadilly/api/http"
	"io/ioutil"
)

type ginReader struct {
	*gin.Context
}

func (g ginReader) GetParameter(key string) string {
	return g.Param(key)
}

func (g ginReader) GetFormData(key string) (string, bool) {
	return g.GetPostForm(key)
}

func (g ginReader) GetHeaders() map[string][]string {
	return g.Context.Request.Header
}

func (g ginReader) ReadBody(obj interface{}) error {
	asd, err := ioutil.ReadAll(g.Request.Body)
	fmt.Println(string(asd))
	fmt.Println(err)
	return g.ShouldBindJSON(obj)
}

func (g ginReader) GetUrl() string {
	return g.Request.URL.String()
}

func CreateReader(ctx *gin.Context) http.Reader {
	return &ginReader{ctx}
}
