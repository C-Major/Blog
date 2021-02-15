package main

import (
	"os"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	common.InitLog()

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config, err := common.GetConfig(rootDir, "conf")
	if err != nil {
		panic(err)
	}

	if err := caller.InitCaller(config); err != nil {
		panic(err)
	}

	router = gin.Default()
	router.LoadHTMLGlob("template/*")

	initializeRoutes()

	router.Run()

}
