package structures

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/crontab"
)

type ParserContext struct {
	Config  *Config
	Gin     *gin.Engine
	Docker  *client.Client
	Cache   *Cache
	Crontab *crontab.Crontab
	Tasks   map[string]*Task
}
