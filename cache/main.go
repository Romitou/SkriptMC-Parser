package cache

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/romitou/skriptmc-parser/structures"
	"log"
	"time"
)

func InitCache(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Initializing cache...")
	ctx.Cache = &structures.Cache{}
	RefreshContainers(ctx)
	log.Println("[Skript-MC] Cache initialized.")
}

func RefreshContainers(ctx *structures.ParserContext) {
	// Clear containers cache
	ctx.Cache.Containers = []types.Container{}

	// Get list of all containers
	containers, err := ctx.Docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return
	}

	// Check if the container is a Skript-MC parser
	for i := range containers {
		container := containers[i]

		// TODO: check the container version
		_, ok := container.Labels["skriptmc.parser.engine-version"]
		if !ok {
			continue
		}

		// Add the container to cache
		ctx.Cache.Containers = append(ctx.Cache.Containers, container)
	}

	ctx.Cache.Updated = time.Now()
}
