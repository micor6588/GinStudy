package main

import (
	"gin-vue-samples/chapter06-02/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
