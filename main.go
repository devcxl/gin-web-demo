package main

import (
	"github.com/gin-gonic/gin"
	"log"
	. "main/config"
	. "main/database"
	. "main/handler"
)

func main() {
	defer SqlDB.Close()
	server := gin.Default()
	server.Use(ServerErrorHandler)
	server.NoRoute(NotFoundHandler)
	server.Use(AuthHandler)
	initRouter(server)
	err := server.Run(":" + Conf.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

}
