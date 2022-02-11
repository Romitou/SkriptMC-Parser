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
	ctx.Gin = gin.Default()

	ctx.Gin.ForwardedByClientIP = true

	// Setup middlewares
	middlewares.RateLimiter(ctx)

	// Register routes
	ctx.Gin.GET("/parsers", createHandler(ctx, handlers.Parsers))
	ctx.Gin.POST("/parse", createHandler(ctx, handlers.Parse))

	err := http.ListenAndServe(ctx.Config.HTTPAddress(), ctx.Gin)
	if err != nil {
		log.Fatal("[Skript-MC] Cannot start web server: ", err)
	}

	log.Println("[Skript-MC] Web server successfully started.")
}

func createHandler(ctx *structures.ParserContext, router func(req structures.HTTPRequest)) func(ginCtx *gin.Context) {
	return func(ginCtx *gin.Context) {
		router(structures.HTTPRequest{
			Gin: ginCtx,
			Ctx: ctx,
		})
	}
}
