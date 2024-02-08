package models

import (
	"encoding/json"
	"net/http"
)

// структура таблицы
type GoodResponse struct {
	Message  string `json:"message"        example:"message"`
	Code     int    `json:"code"           example:"status"`
	Affected int    `json:"affected_rows"  example:"0"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func BadResponseSend(w http.ResponseWriter, message string, code int) {
	badResponse := ErrorResponse{
		Message: message,
		Code:    code,
	}
	json.NewEncoder(w).Encode(badResponse)
}

func GoodResponseSend(w http.ResponseWriter, message string, affectedRow int) {
	errResp := GoodResponse{
		Message:  message,
		Code:     200,
		Affected: affectedRow,
	}
	json.NewEncoder(w).Encode(errResp)
}
