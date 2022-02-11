package http

import (
	"github.com/gin-gonic/gin"
	"github.com/romitou/skriptmc-parser/http/handlers"
	"github.com/romitou/skriptmc-parser/http/middlewares"
	"github.com/romitou/skriptmc-parser/structures"
	"log"
	"net/http"
)

func StartHTTP(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Starting web server...")
	r := gin.Default()

	// Register routes
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

	// Setup middlewares
	middlewares.RateLimiter(ctx)

	err := http.ListenAndServe(ctx.Config.HTTPAddress(), r)
	if err != nil {
		log.Fatal("[Skript-MC] Cannot start web server: ", err)
	}

	log.Println("[Skript-MC] Web server successfully started.")
	ctx.Gin = r
}
