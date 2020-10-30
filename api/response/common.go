package response

type BasicResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	*BasicResponse
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
	*BasicResponse
}

func NewErrorResponse(code int, msg string) *ErrorResponse {
	return &ErrorResponse{Code: code, BasicResponse: &BasicResponse{Status: "error", Message: msg}}
}

func NewFailResponse(msg string) *BasicResponse {
	return &BasicResponse{Status: "fail", Message: msg}
}

func NewSuccessResponse(data interface{}, msg string) *SuccessResponse {
	return &SuccessResponse{Data: data, BasicResponse: &BasicResponse{Status: "success", Message: msg}}
}
