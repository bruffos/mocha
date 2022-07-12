package test

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vitorsalgado/mocha"
	"github.com/vitorsalgado/mocha/expect"
	"github.com/vitorsalgado/mocha/reply"
	"github.com/vitorsalgado/mocha/x/headers"
	"github.com/vitorsalgado/mocha/x/mimetypes"
)

func TestFormUrlEncoded(t *testing.T) {
	m := mocha.New(t)
	m.Start()

	scoped := m.Mock(mocha.Post(expect.URLPath("/test")).
		FormField("var1", expect.ToEqual("dev")).
		FormField("var2", expect.ToContain("q")).
		Reply(reply.OK()))

	data := url.Values{}
	data.Set("var1", "dev")
	data.Set("var2", "qa")

	req, _ := http.NewRequest(http.MethodPost, m.URL()+"/test", strings.NewReader(data.Encode()))
	req.Header.Add("test", "hello")
	req.Header.Add(headers.ContentType, mimetypes.FormURLEncoded)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	assert.True(t, scoped.Called())
}
