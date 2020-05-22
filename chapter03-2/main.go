package main

import (
	"gin-vue-samples/chapter03-2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Route(engine)
	_ = engine.Run() // listen and serve on 0.0.0.0:8080
}
