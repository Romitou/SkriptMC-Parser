package structures

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/crontab"
)

type ParserContext struct {
	HTTP    *gin.Engine
	Docker  *client.Client
	Cache   *Cache
	Crontab *crontab.Crontab
	Tasks   map[string]*Task
}
