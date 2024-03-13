package controllers

import (
	"net/http"

	"github.com/DivyaMaddipudi/receipt-processor-challenge/helpers"
	"github.com/DivyaMaddipudi/receipt-processor-challenge/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Controller to get all receipts
func GetAllReceipts(context *gin.Context)  {
	context.IndentedJSON(http.StatusOK, models.Receipts)
}

//Controller to add a new receipt
func AddReceipt(context *gin.Context)  {
	var newReceipt models.Receipt
	if err := context.BindJSON(&newReceipt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt"})
		return 
	}

	isEmptyBody := helpers.ValidateReceipt(newReceipt)

	if isEmptyBody {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format of receipt"})
		return 
	}

	newReceipt.ID = uuid.New().String()

	models.Receipts = append(models.Receipts, newReceipt)
	
	// Return the ID of the newly added receipt
	context.JSON(http.StatusCreated, gin.H{"id": newReceipt.ID})
}

//Contoller to calculate receipt points based on receipt id
func GetReceiptPointsById(context *gin.Context)  {
	id := context.Param("id")

	receipt, index, err := helpers.GetReceiptById(id, models.Receipts)

	if err != "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message" : err})
		return
	}

	if receipt.Points != 0 {
		context.JSON(http.StatusCreated, gin.H{"points": receipt.Points})
		return
	}

	points, err := helpers.CalculatePointsByReceipt(receipt)

	if err != "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message" : err})
		return
	}
	models.Receipts[index].Points = points
	
	context.JSON(http.StatusCreated, gin.H{"points": points})
}