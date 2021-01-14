package main

import (
	"github.com/gin-gonic/gin"
	. "main/apis"
)

func initRouter(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/version", AddUserApi)
		user := v1.Group("/user")
		{
			user.POST("", AddUserApi)
			user.GET("/hello", IndexApi)
		}
		v1.GET("/persons", GetPersonsApi)
		person := v1.Group("/person")
		{
			person.POST("", AddPersonApi)
			person.GET("/:id", GetPersonApi)
			person.PUT("/:id", ModPersonApi)
			person.DELETE("/:id", DelPersonApi)
		}
	}
}

func initRouterV2(engine *gin.Engine) {

}
