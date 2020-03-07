package response

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpStatus int, success bool, message string, data interface{}) {
	g.C.JSON(httpStatus, Response{
		Success: success,
		Message: message,
		Data:    data,
	})

	return
}
