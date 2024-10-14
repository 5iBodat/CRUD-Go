package Response

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string) *BaseResponse {
	return &BaseResponse{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}

// ToJSON converts the BaseResponse to JSON format
func (br *BaseResponse) ToJSON() ([]byte, error) {
	return json.Marshal(br)
}

// FromJSON converts the BaseResponse
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := SuccessResponse(message, data)
	WriteJSONResponse(w, statusCode, response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	response := ErrorResponse(message)
	WriteJSONResponse(w, statusCode, response)
}
