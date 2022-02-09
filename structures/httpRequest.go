package structures

import "github.com/gin-gonic/gin"

type HTTPRequest struct {
	Gin *gin.Context
	Ctx *ParserContext
}
