package main

import (
	"github.com/romitou/skriptmc-parser/cache"
	"github.com/romitou/skriptmc-parser/config"
	"github.com/romitou/skriptmc-parser/docker"
	"github.com/romitou/skriptmc-parser/http"
	"github.com/romitou/skriptmc-parser/redis"
	"github.com/romitou/skriptmc-parser/structures"
	"github.com/romitou/skriptmc-parser/tasks"
)

func main() {
	var ctx structures.ParserContext
	config.LoadConfig(&ctx)
	redis.ConnectRedis(&ctx)
	docker.StartDocker(&ctx)
	cache.InitCache(&ctx)
	tasks.InitTasks(&ctx)
	http.StartHTTP(&ctx)
}
