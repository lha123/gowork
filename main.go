package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gowork/api"
	_ "gowork/docs"
	//_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/user")
	api.RegisterUser(v1)

	r.Run(":8089")
}
