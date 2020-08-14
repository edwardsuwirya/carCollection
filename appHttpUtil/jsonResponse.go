package appHttpUtil

import (
	"encoding/json"
	"net/http"
)

type jsonResponder struct {
}

func NewDefaultJSONResponder() HttpResponseBuilder {
	return &jsonResponder{}
}

func (j *jsonResponder) Write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}
	json.NewEncoder(w).Encode(data)
}
func (j *jsonResponder) Data(w http.ResponseWriter, status int, message string, data interface{}) {
	content := DataResponse{Message: message, Data: data}
	j.Write(w, status, content)

}
func (j *jsonResponder) Error(w http.ResponseWriter, status int, error ErrorResponseBuilder) {
	content := ErrorResponse{ErrorID: error.Code(), Message: error.Message()}
	j.Write(w, status, content)
}
