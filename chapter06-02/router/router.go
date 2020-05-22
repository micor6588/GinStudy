package router

import (
	"gin-vue-samples/chapter06-02/controller"

	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	engine.GET("/json", controller.IndexJsonAction)
	engine.GET("/xml", controller.IndexXmlAction)
	engine.GET("/yaml", controller.IndexYamlAction)
	engine.GET("/image", controller.IndexImageAction)
	engine.GET("/file", controller.IndexFileAction)

	engine.GET("/login", controller.LoginGetAction)
	engine.POST("/login", controller.LoginPostAction)
	engine.PUT("/login", controller.LoginPutAction)
}
