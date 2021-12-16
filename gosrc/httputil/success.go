package httputil

import "github.com/gin-gonic/gin"

// NewSuccess example
func NewSuccess(ctx *gin.Context, status int, message string) {
	er := HTTPSuccess{
		Code:    status,
		Message: message,
	}
	ctx.JSON(status, er)
}

// HTTPSuccess example
type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Successfully refreshed cache"`
}
