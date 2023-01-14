package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shomuMatch/goCoverTest/api"
)

type externalResponse struct {
	statusCode int
	value      int
}

func TestApi2(t *testing.T) {
	testCases := []struct {
		request          api.Request
		response         api.Response
		externalResponse externalResponse
	}{
		{
			request: api.Request{
				Value1: 8,
				Value2: 72,
			},
			response: api.Response{Value: 400},
			externalResponse: externalResponse{
				statusCode: http.StatusOK,
				value:      4,
			},
		},
		{
			request: api.Request{
				Value1: 0,
				Value2: 14,
			},
			response: api.Response{Value: 101},
			externalResponse: externalResponse{
				statusCode: http.StatusOK,
				value:      101,
			},
		},
	}
	for caseNo, testCase := range testCases {
		t.Run(fmt.Sprint(caseNo), func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.externalResponse.statusCode)
				json.NewEncoder(w).Encode(api.Response{Value: testCase.externalResponse.value})
			})
			l, _ := net.Listen("tcp", ":8008")
			ts := httptest.Server{
				Listener: l,
				Config:   &http.Server{Handler: handler},
			}
			ts.Start()
			defer ts.Close()

			bodyJson, _ := json.Marshal(testCase.request)
			url := "http://localhost:8888/api2"
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
