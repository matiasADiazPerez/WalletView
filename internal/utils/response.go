package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"walletview/internal/models"
)

func CreateResponse(message string, payload any, w http.ResponseWriter) {
	response := models.Response{
		Message: message,
		Payload: payload,
	}
	jData, err := json.Marshal(response)
	if err != nil {
		HandleError(NewErrorWrapper("Error creating the response", 0, err), w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}

func NewErrorWrapper(message string, code int, err error) models.ErrorWrapper {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	return models.ErrorWrapper{
		Message: message,
		Error:   err,
		Code:    code,
	}
}

func HandleError(err models.ErrorWrapper, w http.ResponseWriter) {
	response := models.Response{
		Message: err.Message,
		Payload: err.Error.Error(),
	}
	jData, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		jData = []byte(fmt.Sprintf("Unexpected Err: %s", jsonErr.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	w.Write(jData)

}
