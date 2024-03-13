package helpers

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/DivyaMaddipudi/receipt-processor-challenge/models"
)

// CountAlphanumeric - returns the count of total number of alpha numeric characters
func CountAlphaNumeric(input string) int {

	count := 0
	for _, char := range input {
			if unicode.IsLetter(char) || unicode.IsDigit(char) { 
					count++
			}
	}
	return count
}

// IsTotalRoundDollar - returns true if the total is round dollar without cents
func IsTotalRoundDollar(input string) (bool, string) {

	parts := strings.Split(input, ".")
	errorMsg := "Invalid amount value"

	if len(parts) == 1 {
		_, err := strconv.Atoi(input)
		if err != nil {
			return false, errorMsg
		}
		return true, ""
	}

	if len(parts) == 2 && parts[1] == "00" {
		return true, ""
	}
	return false, ""
}

// IsMultipleOfQuarters - returns true if the total is multiple of 0.25
func IsMultipleOfQuarter(input string) (bool, string) {

	totalAmount, err := strconv.ParseFloat(input, 64)
	errorMsg := "Error parsing input"

	if err != nil {
		return false, errorMsg
	}

	totalCents := totalAmount * 100

	return int(totalCents)%25 == 0, ""
}

// CalculatePointsForDescriptionLength - 
// calculates the total points based on the (price * 0.2), if the trimmied item description is divisible by 3  
func CalculatePointsForDescriptionLength(item models.Item) (int, string) {

	trimmedDescription := strings.TrimSpace(item.ShortDescription)
	errorMsg := "Error parsing item description"

	if len(trimmedDescription)%3 == 0 {
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, errorMsg
		}
		points := int(math.Ceil(price * 0.2))
		return points, ""
	}
	return 0, ""
}

// IsOddDay - returns true if day is odd
func IsOddDay(input string) (bool, string) {

	parts := strings.Split(input, "-") 
	errorMsg := "Invalid date"

	if len(parts) != 3 {
		return false, errorMsg
	}
	day, err := strconv.Atoi(parts[2])

	if err != nil {
		return false, errorMsg
	}
	return day%2 != 0, ""
}

//IsTimeBetween2And4PM - returns true if time is between 2PM and 4 PM
func IsTimeBetween2And4PM(input string) (bool, string) {

	parsedTime, err := time.Parse("15:04", input)
	errorMsg := "Invalid Purchase Time"

	if err != nil {
		return false, errorMsg
	}

	// Define time objects for 2:00 PM and 4:00 PM on the same day as the input time
	startTime := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 14, 0, 0, 0, parsedTime.Location())
	endTime := time.Date(parsedTime.Year(), parsedTime.Month(), parsedTime.Day(), 16, 0, 0, 0, parsedTime.Location())

	return parsedTime.After(startTime) && parsedTime.Before(endTime), ""
}