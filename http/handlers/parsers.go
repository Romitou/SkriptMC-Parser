package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/romitou/skriptmc-parser/structures"
)

func Parsers(req structures.HTTPRequest) {
	containers := req.Ctx.Cache.Containers
	updated := req.Ctx.Cache.Updated

	containerInfo := []structures.ContainerInfo{}
	for i := range containers {
		container := containers[i]
		version := container.Labels["skriptmc.parser.engine-version"]
		containerInfo = append(containerInfo, structures.ContainerInfo{
			Name:    container.Names[0],
			Version: version,
			State:   container.State,
			Status:  container.Status,
		})
	}

	req.Gin.JSON(200, gin.H{"lastUpdated": updated, "parsers": containerInfo})
}
