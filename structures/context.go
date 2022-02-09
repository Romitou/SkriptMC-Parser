package structures

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type ParserContext struct {
	HTTP   *gin.Engine
	Docker *client.Client
	Cache  *Cache
}
