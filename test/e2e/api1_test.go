package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/shomuMatch/goCoverTest/api"
)

func TestApi1(t *testing.T) {
	testCases := []struct {
		request  api.Request
		response api.Response
	}{
		{
			request: api.Request{
				Value1: 8,
				Value2: 72,
			},
			response: api.Response{Value: 9},
		},
		{
			request: api.Request{
				Value1: 0,
				Value2: 14,
			},
			response: api.Response{Value: 1400},
		},
	}
	for caseNo, testCase := range testCases {
		t.Run(fmt.Sprint(caseNo), func(t *testing.T) {
			bodyJson, _ := json.Marshal(testCase.request)
			url := "http://localhost:8888/api1"
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
			client := new(http.Client)
			res, _ := client.Do(req)
			if res.StatusCode != 200 {
				t.Errorf("response status is %v", res.StatusCode)
			}
			resBody, _ := io.ReadAll(res.Body)
			var response api.Response
			json.Unmarshal(resBody, &response)
			if testCase.response.Value != response.Value {
				t.Errorf("expect: %v, actual: %v", testCase.response.Value, response.Value)
			}
		})
	}
}
