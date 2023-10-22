package entity

type BaseResponse struct {
	Status  int                `json:"status"`
	Type    string             `json:"type"`
	Message string             `json:"message"`
	Detail  BaseResponseDetail `json:"detail"`
}

type BaseResponseDetail struct {
	Datas  any                `json:"datas,omitempty"`
	Errors []ErrorFieldDetail `json:"errors,omitempty"`
}

type ErrorFieldDetail struct {
	Field  string `json:"field"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
}
