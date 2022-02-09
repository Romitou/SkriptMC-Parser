package structures

import (
	"github.com/docker/docker/api/types"
	"time"
)

type Cache struct {
	Containers []types.Container
	Updated    time.Time
}
