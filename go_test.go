package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFail1(t *testing.T) {
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
