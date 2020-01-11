package handlers


type FailedResponse struct {
	ErrMsg string `json:"err_msg"`
	ErrCode int `json:"err_code"`
}