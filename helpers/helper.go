package helpers

import (
	"github.com/DivyaMaddipudi/receipt-processor-challenge/models"
)

func ValidateReceipt(receipts models.Receipt) bool {

	return receipts.Retailer == "" ||
			receipts.PurchaseDate == "" ||
			receipts.PurchaseTime == "" ||
			len(receipts.Items) == 0 ||
			receipts.Total == ""
}

func GetReceiptById(id string, receipts []models.Receipt) (models.Receipt, int, string) {
	var receipt models.Receipt
	errorMsg := "Invalid Receipt Id"
	for index, receipt := range receipts {
		if receipt.ID == id {
			return receipt, index,  ""
		}
	}
	return receipt, -1, errorMsg
}

func CalculatePointsByReceipt(receipt models.Receipt) (int, string) {
	points := 0
	
	//Adding points based on retailer name
	points += CountAlphaNumeric(receipt.Retailer)

	//Checking the total is rounded or not
	isTotalRoundedDollar, err := IsTotalRoundDollar(receipt.Total)
	if err != "" {
        return 0, err 
    }
	if isTotalRoundedDollar {
		points += 50
	}

	//Checking the total is multiple by 0.25 or not
	isMultipleOfQuater, err := IsMultipleOfQuarter(receipt.Total)
	if err != "" {
        return 0, err 
    }
	if isMultipleOfQuater {
		points += 25
	}

	//Adding points based on the total items
	points += 5 * (len((receipt.Items))/2)

	//Calculating points based on the length of the item description
	for _, item := range receipt.Items {
		itemPoints, err := CalculatePointsForDescriptionLength(item)
		if err != "" {
			return 0, err
		}
		points += itemPoints
	}

	//Adding points based on odd day
	isOddDay, err := IsOddDay(receipt.PurchaseDate)
	if err != "" {
        return 0, err 
    }
	if isOddDay {
		points += 6
	}

	//Adding points when the purchased time is between 2PM and 4PM
	isTimeBetween2And4PM, err := IsTimeBetween2And4PM(receipt.PurchaseTime)
	if err != "" {
        return 0, err 
    }
	if isTimeBetween2And4PM {
		points += 10
	}

	return points, ""
}