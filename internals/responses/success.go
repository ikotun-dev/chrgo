package responses

import (
	"encoding/json"
	"net/http"
)

type HTTPResp struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func RequestSuccess(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	successReponse := HTTPResp{Message: message, Code: code}
	json.NewEncoder(w).Encode(successReponse)

}
