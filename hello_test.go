package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

type JsonResponse struct {
	Foo string
	Bar string
}

func TestMain(m *testing.M) {
	fmt.Println("main")
	data := url.Values{}
	data.Set("name", "foo")
	req, _ := http.NewRequest("POST", "http://localhost:8090/hello", strings.NewReader(data.Encode()))
	httpmock.ConnectionFailure(req)

	a := JsonResponse{Foo: "foo test", Bar: "bar test"}
	httpmock.RegisterResponder("GET", "http://localhost:8090/hello",
		func(req *http.Request) (*http.Response, error) {
			fmt.Println("call")
			return httpmock.NewJsonResponse(200, a)
		})

	httpmock.RegisterResponder("POST", "http://localhost:8090/hello",
		func(req *http.Request) (*http.Response, error) {
			fmt.Printf("call %v\n", req)
			return httpmock.NewJsonResponse(200, a)
		})

	os.Exit(m.Run())
}

func TestHttp(t *testing.T) {
	httpmock.Activate()
	t.Cleanup(func() { httpmock.DeactivateAndReset() })

	if resp, err := http.Get("http://localhost:8090/hello"); err == nil {
		defer resp.Body.Close()
		jsonResponse := &JsonResponse{}
		json.NewDecoder(resp.Body).Decode(jsonResponse)
		fmt.Printf("Bar : %v\nFoo : %v\n", jsonResponse.Bar, jsonResponse.Foo)
	} else {
		fmt.Printf("%v err\n", err)
	}

}

func TestFail(t *testing.T) {
	httpmock.Activate()

	data := url.Values{}
	data.Set("name", "foo")

	if resp, err := http.Post("http://localhost:8090/hello", "json/application", strings.NewReader(data.Encode())); err == nil {
		defer resp.Body.Close()
		jsonResponse := &JsonResponse{}
		json.NewDecoder(resp.Body).Decode(jsonResponse)
		fmt.Printf("Bar : %v\nFoo : %v\n", jsonResponse.Bar, jsonResponse.Foo)
	} else {
		fmt.Printf("%v err\n", err)
	}
	// httpmock.ConnectionFailure()
}
