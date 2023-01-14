package path2

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shomuMatch/goCoverTest/api"
)

func Api2(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var request api.Request
	json.Unmarshal(body, &request)

	url := "http://localhost:8008/external/api"
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	client := new(http.Client)
	res, _ := client.Do(req)
	if res.StatusCode != 200 {
		json.NewEncoder(w).Encode(api.Response{Value: res.StatusCode})
		return
	}
	resBody, _ := io.ReadAll(res.Body)
	var externalResponse api.Response
	json.Unmarshal(resBody, &externalResponse)

	var val int
	if externalResponse.Value%2 == 0 {
		val = externalResponse.Value * 100
	} else {
		val = externalResponse.Value
	}
	json.NewEncoder(w).Encode(api.Response{Value: val})

}
