package path1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shomuMatch/goCoverTest/api"
)

func Api1(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var request api.Request
	json.Unmarshal(body, &request)

	val := 44
	if request.Value1 != 0 {
		val = request.Value2 / request.Value1
	} else if request.Value2 != 0 {
		val = request.Value2 * 100
	} else {
		val = 5000000000000000
	}
	json.NewEncoder(w).Encode(api.Response{Value: val})
}
