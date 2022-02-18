package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/romitou/skriptmc-parser/structures"
)

func Cors(ctx *structures.ParserContext) {
	ctx.Gin.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})
}
