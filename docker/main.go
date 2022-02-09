package docker

import (
	"github.com/docker/docker/client"
	"github.com/romitou/skriptmc-parser/structures"
	"log"
	"runtime"
)

func StartDocker(ctx *structures.ParserContext) {
	log.Println("[Skript-MC] Creating docker client...")

	var host string
	if runtime.GOOS == "windows" { // Support for Windows (for development purposes only)
		host = "npipe:////.//pipe//docker_engine"
	} else {
		host = "unix:///var/run/docker.sock"
	}
	cli, err := client.NewClientWithOpts(client.WithHost(host))
	if err != nil {
		log.Fatal("[Skript-MC] Cannot create docker client: ", err)
	}

	log.Println("[Skript-MC] Docker client successfully created.")
	ctx.Docker = cli
}
