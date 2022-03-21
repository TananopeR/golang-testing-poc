package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("init http_test")
}

type JsonResponse struct {
	Foo string
	Bar string
}

var a JsonResponse

func TestMain(m *testing.M) {
	a = JsonResponse{Foo: "foo test", Bar: "bar test"}
	fmt.Println("main http_test")

	httpmock.RegisterResponder("GET", "http://localhost:8090/hello",
		func(req *http.Request) (*http.Response, error) {
			fmt.Println("call GET")
			return httpmock.NewJsonResponse(200, a)
		})

	httpmock.RegisterResponder("POST", "http://localhost:8090/hello",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewJsonResponse(200, a)
		})

	run := m.Run()
	fmt.Println("main finished")
	os.Exit(run)
}

func TestHttp(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(func() { httpmock.DeactivateAndReset() })

	jsonResponse := &JsonResponse{}

	if resp, err := http.Get("http://localhost:8090/hello"); err == nil {
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(jsonResponse)
		fmt.Printf("Bar : %v\nFoo : %v\n", jsonResponse.Bar, jsonResponse.Foo)
	} else {
		fmt.Printf("%v err\n", err)
	}

	assert.EqualValues(t, jsonResponse.Bar, a.Bar)
	assert.EqualValues(t, jsonResponse.Foo, a.Foo)
	assert.EqualValues(t, 1, httpmock.GetCallCountInfo()["GET http://localhost:8090/hello"])
	assert.EqualValues(t, 0, httpmock.GetCallCountInfo()["POST http://localhost:8090/hello"])
}

func TestA(t *testing.T) {
	httpmock.Activate()
	b := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("handler %v %v\n", w, r.URL)
	})
	u := url.Values{}
	u.Add("test", "abc")
	ass := assert.HTTPBody(b, "GET", "http://localhost:8090/hello", u)
	fmt.Println(ass)

}
