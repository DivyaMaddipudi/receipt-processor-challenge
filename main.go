package main

import (
	"fmt"

	"github.com/DivyaMaddipudi/receipt-processor-challenge/router"
	"github.com/gin-gonic/gin"
)


func main()  {
	fmt.Println("Getting started")
	routes := gin.Default()
	router.ImportRoutes(routes)
	routes.Run("localhost:8000")
}