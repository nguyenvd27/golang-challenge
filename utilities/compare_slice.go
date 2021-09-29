package utilities

import (
	"golang-coding-challenge/transfers"
)

func CompareTransactionJsonSlice(slice1 []transfers.TransactionJson, slice2 []transfers.TransactionJson) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, _ := range slice1 {
		if (slice1[i].Id != slice2[i].Id) || (slice1[i].AccountID != slice2[i].AccountID) ||
			(slice1[i].Amount != slice2[i].Amount) || (slice1[i].Bank != slice2[i].Bank) ||
			(slice1[i].TransactionType != slice2[i].TransactionType) || (slice1[i].CreatedAt != slice2[i].CreatedAt) {
			return false
		}
	}

	return true
}

// func CompareTransactionSlice(slice1 []entities.Transaction, slice2 []entities.Transaction) bool {
// 	if len(slice1) != len(slice2) {
// 		return false
// 	}

// 	for i, _ := range slice1 {
// 		if (slice1[i].Model != slice2[i].Model) || (slice1[i].AccountID != slice2[i].AccountID) || (slice1[i].Amount != slice2[i].Amount) ||
// 			(slice1[i].TransactionType != slice2[i].TransactionType) || (slice1[i].CreatedAt != slice2[i].CreatedAt) {
// 			return false
// 		}
// 	}

// 	return true
// }
