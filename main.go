package main

import (
	"github.com/romitou/skriptmc-parser/cache"
	"github.com/romitou/skriptmc-parser/docker"
	"github.com/romitou/skriptmc-parser/http"
	"github.com/romitou/skriptmc-parser/structures"
)

func main() {
	var ctx structures.ParserContext
	docker.StartDocker(&ctx)
	cache.InitCache(&ctx)
	http.StartHTTP(&ctx)
}
