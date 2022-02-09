package handlers

import (
	"github.com/romitou/skriptmc-parser/structures"
	"io/ioutil"
	"math/rand"
	"time"
)

func Parse(req structures.HTTPRequest) {
	body, err := ioutil.ReadAll(req.Gin.Request.Body)
	if err != nil {
		req.Gin.Status(400)
		return
	}
	containers := req.Ctx.Cache.Containers

	rand.Seed(time.Now().Unix())

	container := containers[rand.Intn(len(containers))]

	// TODO: make request to the parser with user script
}
