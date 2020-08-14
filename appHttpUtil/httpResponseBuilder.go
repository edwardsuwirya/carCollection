package appHttpUtil

import "net/http"

type HttpResponseBuilder interface {
	Data(w http.ResponseWriter, status int, message string, data interface{})
	Error(w http.ResponseWriter, status int, error ErrorResponseBuilder)
}

type ErrorResponse struct {
	ErrorID int    `json:"errorId"`
	Message string `json:"message"`
}

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
