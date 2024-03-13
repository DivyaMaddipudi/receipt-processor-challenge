package router

import (
	"github.com/DivyaMaddipudi/receipt-processor-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func ImportRoutes(routerEngine *gin.Engine)  {
	//Route to get all the receipts
	routerEngine.GET("/receipts", controllers.GetAllReceipts)

	//Route to add new receipt
	routerEngine.POST("/receipts/process", controllers.AddReceipt)

	//Route to calculate points based on the receipt id
	routerEngine.GET("/receipts/:id/points", controllers.GetReceiptPointsById)
}