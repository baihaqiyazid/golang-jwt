package web

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ToWebResponse(code int, status string, message interface{}, data interface{}) WebResponse {
	webResponse := WebResponse{
		Code:   code,
		Status: status,
		Message: message,
		Data:   data,
	}

	return webResponse
}
