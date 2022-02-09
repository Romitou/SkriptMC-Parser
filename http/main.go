package http

import (
	"github.com/gin-gonic/gin"
	"github.com/romitou/skriptmc-parser/http/handlers"
	"github.com/romitou/skriptmc-parser/structures"
	"log"
)

func StartHTTP(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Starting web server...")
	r := gin.Default()

	r.GET("/parsers", func(ginCtx *gin.Context) {
		handlers.Parsers(structures.HTTPRequest{
			Gin: ginCtx,
			Ctx: ctx,
		})
	})

	r.POST("/parse", func(ginCtx *gin.Context) {
		handlers.Parse(structures.HTTPRequest{
			Gin: ginCtx,
			Ctx: ctx,
		})
	})

	err := r.Run("0.0.0.0:8000")
	if err != nil {
		log.Fatal("[Skript-MC] Cannot start web server: ", err)
	}

	log.Println("[Skript-MC] Web server successfully started.")
	ctx.HTTP = r
}
