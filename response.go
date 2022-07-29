package manygo

type BoolSuccessResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	ErrorCode int         `json:"error_code,omitempty"`
}
