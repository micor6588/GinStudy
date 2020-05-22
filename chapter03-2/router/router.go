package router

import (
	"gin-vue-samples/chapter03-2/controller"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	engine.GET("/", controller.IndexGetAction)
	engine.POST("/", controller.IndexPostAction)
	engine.PUT("/", controller.IndexPutAction)
	engine.DELETE("/", controller.IndexDeleteAction)
}
