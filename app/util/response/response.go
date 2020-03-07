package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(success bool, message string, data interface{}) {
	g.C.JSON(http.StatusOK, response{
		Success: success,
		Message: message,
		Data:    data,
	})

	return
}
