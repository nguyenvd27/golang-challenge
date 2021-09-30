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

func CompareTransactionJson(trans1 *transfers.TransactionJson, trans2 *transfers.TransactionJson) bool {
	if trans1 == nil && trans2 == nil {
		return true
	}

	if (trans1 == nil && trans2 != nil) || (trans1 != nil && trans2 == nil) {
		return false
	}

	if (trans1.Id != trans2.Id) || (trans1.AccountID != trans2.AccountID) ||
		(trans1.Amount != trans2.Amount) || (trans1.Bank != trans2.Bank) ||
		(trans1.TransactionType != trans2.TransactionType) || (trans1.CreatedAt != trans2.CreatedAt) {
		return false
	}
	return true
}
