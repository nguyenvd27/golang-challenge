package usecases

import (
	"fmt"
	"golang-coding-challenge/models/entities"
	"golang-coding-challenge/repositories/mock"
	"golang-coding-challenge/transfers"
	"golang-coding-challenge/utilities"
	"reflect"
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

		if utilities.CompareTransactionJsonSlice(result, expected_result) == false || err.Error() != expectedErr.Error() {
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

		if utilities.CompareTransactionJsonSlice(result, expected_result) == false || err.Error() != expectedErr.Error() {
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

func Test_transactionUsecase_CreateTransaction(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		// prepare input
		userID := 1
		reqBody := CreateTransactionRequest{
			AccountID:       1,
			Amount:          1234.56,
			TransactionType: "deposit",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, reqBody.AccountID).Return(
			[]int{},
			fmt.Errorf("not found account"),
		)

		// output
		var expected_result *transfers.TransactionJson = nil
		expectedErr := fmt.Errorf("not found account")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: nil,
		}

		result, err := uc.CreateTransaction(reqBody, userID)

		if utilities.CompareTransactionJson(result, expected_result) == false || err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		// prepare input
		userID := 1
		reqBody := CreateTransactionRequest{
			AccountID:       1,
			Amount:          1234.56,
			TransactionType: "deposit",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, reqBody.AccountID).Return(
			[]int{},
			nil,
		)

		// output
		var expected_result *transfers.TransactionJson = nil
		expectedErr := fmt.Errorf("not found user account")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: nil,
		}

		result, err := uc.CreateTransaction(reqBody, userID)

		if utilities.CompareTransactionJson(result, expected_result) == false || err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		// prepare input
		userID := 1
		reqBody := CreateTransactionRequest{
			AccountID:       1,
			Amount:          1234.56,
			TransactionType: "deposit",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, reqBody.AccountID).Return(
			[]int{1},
			nil,
		)

		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		transactionRepo.EXPECT().CreateTransaction(entities.Transaction{
			AccountID:       uint(reqBody.AccountID),
			Amount:          reqBody.Amount,
			TransactionType: reqBody.TransactionType,
		}).Return(
			entities.Transaction{},
			fmt.Errorf("can not create transaction"),
		)

		// output
		var expected_result *transfers.TransactionJson = nil
		expectedErr := fmt.Errorf("can not create transaction")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.CreateTransaction(reqBody, userID)

		if utilities.CompareTransactionJson(result, expected_result) == false || err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", result, expected_result)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		// prepare input
		userID := 1
		reqBody := CreateTransactionRequest{
			AccountID:       1,
			Amount:          1234.56,
			TransactionType: "deposit",
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, reqBody.AccountID).Return(
			[]int{1},
			nil,
		)

		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		timeTmp := time.Now()
		transactionRepo.EXPECT().CreateTransaction(entities.Transaction{
			AccountID:       uint(reqBody.AccountID),
			Amount:          reqBody.Amount,
			TransactionType: reqBody.TransactionType,
		}).Return(
			entities.Transaction{
				Model: gorm.Model{
					ID:        20,
					CreatedAt: timeTmp,
					UpdatedAt: timeTmp,
				},
				Amount:          reqBody.Amount,
				TransactionType: reqBody.TransactionType,
				AccountID:       uint(reqBody.AccountID),
				Account: entities.Account{
					Bank: "VCB",
				},
			},
			nil,
		)

		// output
		var expected_result *transfers.TransactionJson = &transfers.TransactionJson{
			Id:              20,
			AccountID:       1,
			Amount:          1234.56,
			Bank:            "VCB",
			TransactionType: "deposit",
			CreatedAt:       timeTmp.String(),
		}
		var expectedErr error = nil

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.CreateTransaction(reqBody, userID)

		if utilities.CompareTransactionJson(result, expected_result) == false || err != expectedErr {
			t.Errorf("error mismatch [%v] [%v]", result, expected_result)
		}
	})
}

func TestNewTransactionUsecase(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// input
		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)

		// output
		expectedResult := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result := NewTransactionUsecase(accountRepo, transactionRepo)

		if reflect.DeepEqual(result, expectedResult) == false {
			t.Errorf("error mismatch [%v] [%v]", result, expectedResult)
		}
	})
}

func Test_transactionUsecase_GetTransactionsV2(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1
		page := 1
		size := 3

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{},
			fmt.Errorf("not found account"),
		)

		// output
		var expectedResult *transfers.PaginateData = nil
		expectedErr := fmt.Errorf("not found account")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: nil,
		}

		result, err := uc.GetTransactionsV2(userID, accountID, page, size)

		if reflect.DeepEqual(result, expectedResult) == false || err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1
		page := 1
		size := 3

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{1},
			nil,
		)

		account_ids := []int{1}
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		transactionRepo.EXPECT().GetTransactionsWithPaginate(account_ids, page, size).Return(
			[]entities.Transaction{},
			int64(0),
			fmt.Errorf("not found transaction"),
		)

		// output
		var expectedResult *transfers.PaginateData = nil
		expectedErr := fmt.Errorf("not found transaction")

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactionsV2(userID, accountID, page, size)

		if reflect.DeepEqual(result, expectedResult) == false || err.Error() != expectedErr.Error() {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1
		page := 1
		size := 3

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		accountRepo := mock.NewMockAccountRepo(mockCtrl)
		accountRepo.EXPECT().GetAccountIDsByUserIDAccountID(userID, accountID).Return(
			[]int{1},
			nil,
		)

		account_ids := []int{1}
		transactionRepo := mock.NewMockTransactionRepo(mockCtrl)
		transactionRepo.EXPECT().GetTransactionsWithPaginate(account_ids, page, size).Return(
			[]entities.Transaction{},
			int64(0),
			nil,
		)

		// output
		var expectedResult *transfers.PaginateData = &transfers.PaginateData{
			Data:  []transfers.TransactionJson{},
			Total: int(0),
			Page:  page,
		}
		var expectedErr error = nil

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactionsV2(userID, accountID, page, size)

		if reflect.DeepEqual(result, expectedResult) == false || err != expectedErr {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})

	t.Run("case 4", func(t *testing.T) {
		// prepare input
		userID := 1
		accountID := 1
		page := 1
		size := 2

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
		transactionRepo.EXPECT().GetTransactionsWithPaginate(account_ids, page, size).Return(
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
				{
					Model: gorm.Model{
						ID:        2,
						CreatedAt: timeTmp,
						UpdatedAt: timeTmp,
					},
					Amount:          2345.67,
					TransactionType: "withdraw",
					AccountID:       1,
					Account: entities.Account{
						Bank: "VCB",
					},
				},
			},
			int64(5),
			nil,
		)

		// output
		var expectedResult *transfers.PaginateData = &transfers.PaginateData{
			Data: []transfers.TransactionJson{
				{
					Id:              1,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
				{
					Id:              2,
					AccountID:       1,
					Amount:          2345.67,
					Bank:            "VCB",
					TransactionType: "withdraw",
					CreatedAt:       timeTmp.String(),
				},
			},
			Total: int(5),
			Page:  page,
		}
		var expectedErr error = nil

		uc := &transactionUsecase{
			accountRepo:     accountRepo,
			transactionRepo: transactionRepo,
		}

		result, err := uc.GetTransactionsV2(userID, accountID, page, size)

		if reflect.DeepEqual(result, expectedResult) == false || err != expectedErr {
			t.Errorf("error mismatch [%v] [%v]", err, expectedErr)
		}
	})
}
