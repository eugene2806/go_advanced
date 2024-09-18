package helpers

import "github.com/gin-gonic/gin"

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func RespondJSON(w *gin.Context, status int, data interface{}) {
	var msg Message

	msg.StatusCode = status
	//meta ...
	msg.Data = data
	w.JSON(status, msg)

}
