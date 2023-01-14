package api

type Request struct {
	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
}

type Response struct {
	Value int `json:"value"`
}
