package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/romitou/skriptmc-parser/structures"
	"log"
)

func ConnectRedis(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Connecting to the Redis server...")
	ctx.Redis = redis.NewClient(&redis.Options{
		Addr:     ctx.Config.Redis.Address,
		Username: ctx.Config.Redis.Username,
		Password: ctx.Config.Redis.Password,
		DB:       ctx.Config.Redis.DB,
	})
	log.Println("[Skript-MC] Redis successfully connected.")
}
