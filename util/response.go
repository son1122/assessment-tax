package util

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
	} `json:"data"`
}
