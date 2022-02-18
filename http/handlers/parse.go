package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/romitou/skriptmc-parser/structures"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func Parse(req structures.HTTPRequest) {
	req.Gin.Header("Access-Control-Allow-Headers", "Content-Type")
	req.Gin.Status(200)
}

func PostParse(req structures.HTTPRequest) {

	requestBody, err := ioutil.ReadAll(req.Gin.Request.Body)
	if err != nil {
		req.Gin.JSON(400, gin.H{
			"code":  2002,
			"error": "Your request body is badly formed or invalid. Correct your request and try again.",
		})
		return
	}

	containers := req.Ctx.Cache.Containers

	if len(containers) == 0 {
		req.Gin.JSON(503, gin.H{
			"code":  2003,
			"error": "No parser is currently able to process your request.",
		})
		return
	}

	rand.Seed(time.Now().Unix())

	container := containers[rand.Intn(len(containers))]

	req.Gin.Header("X-Parser", container.Names[0])

	parserNetwork, ok := container.NetworkSettings.Networks["skmc-parser"]
	if !ok {
		req.Gin.JSON(500, gin.H{
			"code":  2004,
			"error": "A parser was found, but seems to be misconfigured. Retry your request.",
		})
		// TODO: remove the parser from cache to not use it again and re-execute the handler?
		return
	}

	response, err := http.Post("http://"+parserNetwork.IPAddress+"/parse", "text/plain", bytes.NewBuffer(requestBody))
	if err != nil || response.StatusCode != 200 {
		req.Gin.JSON(500, gin.H{
			"code":  2005,
			"error": "Your request has been sent to a parser but it seems to return an error. Try your request again.",
		})
		// TODO: remove the parser from cache to not use it again and re-execute the handler?
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		req.Gin.JSON(500, gin.H{
			"code":  2006,
			"error": "An error occurred while reading the parser response. Try your request again.",
		})
		return
	}

	req.Gin.Data(200, "application/json", body)
}
