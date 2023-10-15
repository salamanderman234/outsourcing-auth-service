package entity

type BaseResponse struct {
	Status  int    `json:"status"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Datas   any    `json:"datas"`
}

type ErrorFieldDetail struct {
	Field  string `json:"field"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}