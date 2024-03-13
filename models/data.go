package models

import "github.com/google/uuid"

var Receipts = []Receipt {
	{
		ID: uuid.New().String(),
		Retailer: "Seven11 Ave123",
		PurchaseDate: "2024-03-10",
		PurchaseTime: "17:30",
		Items: []Item{
			{
				ShortDescription: " Chocolates ",
				Price: "12",
			},
			{
				ShortDescription: "Pizza",
				Price: "19",
			},
		},
		Total: "31.00",
	},

	{
		ID: uuid.New().String(),
		Retailer: "StroesIn123 Ave",
		PurchaseDate: "2024-03-12",
		PurchaseTime: "14:30",
		Items: []Item{
			{
				ShortDescription: "Kitchen items",
				Price: "23",
			},
			{
				ShortDescription: "Items for living room",
				Price: "6.89",
			},
			{
				ShortDescription: "Cookies",
				Price: "3.45",
			},
		},
		Total: "33.34",
	},
}