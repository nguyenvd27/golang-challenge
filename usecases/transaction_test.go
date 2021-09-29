package usecases

import (
	"fmt"
	"golang-coding-challenge/models/entities"
	"golang-coding-challenge/repositories/mock"
	"golang-coding-challenge/transfers"
	"golang-coding-challenge/utilities"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func Test_transactionUsecase_GetTransactions(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{},
			fmt.Errorf("not found account"),
		)

		// output
		var expected_result []transfers.TransactionJson = nil
		expectedErr := fmt.Errorf("not found account")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: nil,
		}

		result, err := uc.GetTransactions(userID, accountID)

		if utilities.CompareTransactionJsonSlice(result, expected_result) == true && err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{1},
			nil,
		)

		account_ids := []int{1}
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		transactionRepo.EXPECT().GetTransactionsByAccountIDs(account_ids).Return(
			[]entities.Transaction{},
			fmt.Errorf("not found transaction"),
		)

		// output
		var expected_result []transfers.TransactionJson = nil
		expectedErr := fmt.Errorf("not found transaction")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactions(userID, accountID)

		if utilities.CompareTransactionJsonSlice(result, expected_result) == true && err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{1},
			nil,
		)

		account_ids := []int{1}
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		transactionRepo.EXPECT().GetTransactionsByAccountIDs(account_ids).Return(
			[]entities.Transaction{},
			nil,
		)

		// output
		expected_result := []transfers.TransactionJson{}
		var expectedErr error = nil

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactions(userID, accountID)

		if utilities.CompareTransactionJsonSlice(result, expected_result) == false || err != expectedErr {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{1},
			nil,
		)

		account_ids := []int{1}
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		timeTmp := time.Date(2021, 9, 23, 11, 12, 16, 0, time.Local)
		transactionRepo.EXPECT().GetTransactionsByAccountIDs(account_ids).Return(
			[]entities.Transaction{
				{
					Model: gorm.Model{
						ID:        1,
						CreatedAt: timeTmp,
						UpdatedAt: timeTmp,
					},
					Amount:          1234.56,
					TransactionType: "deposit",
					AccountID:       1,
					Account: entities.Account{
						Bank: "VCB",
					},
				},
			},
			nil,
		)

		// output
		var expectedErr error = nil
		expected_result := []transfers.TransactionJson{
			{
				Id:              1,
				AccountID:       1,
				Amount:          1234.56,
				Bank:            "VCB",
				TransactionType: "deposit",
				CreatedAt:       timeTmp.String(),
			},
		}

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactions(userID, accountID)

		if utilities.CompareTransactionJsonSlice(result, expected_result) == false || err != expectedErr {
			t.Errorf("error mismatch [%v] [%v]", result, expected_result)
		}
	})
}
