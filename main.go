package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shraddha0602/Go-PostgreSQL/config"
	"github.com/shraddha0602/Go-PostgreSQL/routes"
)

func main() {

	//initialize router
	config.Connect()
	router := gin.Default()

	routes.Routes(router)

	//Listen to the router with port
	log.Fatal(router.Run("localhost:8800"))
}
