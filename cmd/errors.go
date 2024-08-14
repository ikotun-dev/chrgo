// package cmd
// import (
// 	"encoding/json"
// 	"net/http"
// )
//
// //TODO: create all sorts of errors i might want to reuse
//
// type HTTPError struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// }
//
// func BadRequestError(w http.ResponseWriter, message string, code int) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	errorMessage := HTTPError{Message: message, Code: code}
// 	json.NewEncoder(w).Encode(errorMessage)
// }

