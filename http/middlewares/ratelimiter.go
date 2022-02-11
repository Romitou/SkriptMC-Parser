package middlewares

import (
	"github.com/romitou/skriptmc-parser/structures"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
)

func RateLimiter(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Registering limiter middleware...")

	rate, err := limiter.NewRateFromFormatted(ctx.Config.RateLimiter.Limit)
	if err != nil {
		log.Fatal("[Skript-MC] Cannot parse formatted rate limit string: ", err)
	}

	store, err := redis.NewStoreWithOptions(ctx.Redis, limiter.StoreOptions{
		Prefix: ctx.Config.RateLimiter.Prefix,
	})
	if err != nil {
		log.Fatal("[Skript-MC] Cannot create a Redis limiter store: ", err)
	}

	middleware := gin.NewMiddleware(limiter.New(store, rate))
	ctx.Gin.Use(middleware)

	log.Println("[Skript-MC] Ratelimiter middleware successfully registered.")
}
