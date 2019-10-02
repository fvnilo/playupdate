package http

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"error"`
	Message    string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}
