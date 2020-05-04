package schemas
import (
	// Standard library packages
	"net/http"
	"encoding/json"
)

// *****************************************************************************
// Entities schema
// *****************************************************************************

type ApiResponse struct {
	Header      int         `json:"-"`
	Status    	string 		`json:"status"`
	Message   	string 		`json:"message"`
	Data      	interface{} `json:"data"`
}

const SUCCESS = "success"
const ERROR   = "error"

func (response ApiResponse) Success(data interface{}, message string) ApiResponse {
	response.Status  = SUCCESS
	response.Message = message
	response.Data    = data

	if array, ok := data.([]map[string]interface{}); ok {
		if len(array) == 0 {
			response.Data = []string{}
		}
	}

	return response
}

func (response ApiResponse) CustomSuccess(header int, data interface{}, message string) ApiResponse {
	response = response.Success(data, message)
	response.Header = header
	return response
}

func (response ApiResponse) IsSuccess() bool {
	return response.Status == SUCCESS
}

func (response ApiResponse) Error(message string) ApiResponse {
	response.Status = ERROR
	response.Message = message
	return response
}

func (response ApiResponse) CustomError(header int, data interface{}, message string) ApiResponse {
	response = response.Error(message)
	response.Header = header
	response.Data   = data
	return response
}

func (response ApiResponse) IsError() bool {
	return response.Status == ERROR
}

func (response ApiResponse) JsonResponse(w http.ResponseWriter) {

	var header int
	if header == 0 {
		switch response.Status {
		case SUCCESS : header = http.StatusOK
		case ERROR   : header = http.StatusBadRequest
		}
	} else {
		header = response.Header
	}

	respByt, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	w.Write(respByt)
}

func (response ApiResponse) InvalidJsonResponse(w http.ResponseWriter, message string) {
	respByt, _ := json.Marshal(response.Error(message))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(respByt)
}
