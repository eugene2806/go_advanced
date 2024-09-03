package responses

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Error      bool        `json:"error"`
}

func Response200(writer http.ResponseWriter, message interface{}) {
	msg := Message{
		StatusCode: 200,
		Message:    message,
		Error:      false,
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(msg)
}

func Response400(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 400,
		Message:    message,
		Error:      true,
	}

	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(msg)
}

func Response500(writer http.ResponseWriter, message string) {
	msg := Message{
		StatusCode: 500,
		Message:    message,
		Error:      true,
	}

	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(msg)
}
