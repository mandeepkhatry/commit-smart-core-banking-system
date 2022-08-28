package http_response

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, code int, payload interface{}) {
	Response(w, code, map[string]interface{}{"error": payload})
}
func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
