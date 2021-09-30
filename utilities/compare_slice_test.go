package utilities

import (
	"golang-coding-challenge/transfers"
	"testing"
	"time"
)

func TestCompareTransactionJsonSlice(t *testing.T) {
	type args struct {
		slice1 []transfers.TransactionJson
		slice2 []transfers.TransactionJson
	}
	timeTmp := time.Now()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				slice1: []transfers.TransactionJson{
					{
						Id:              1,
						AccountID:       1,
						Amount:          1234.56,
						Bank:            "VCB",
						TransactionType: "deposit",
						CreatedAt:       timeTmp.String(),
					},
				},
				slice2: []transfers.TransactionJson{},
			},
			want: false,
		},
		{
			name: "case 2",
			args: args{
				slice1: []transfers.TransactionJson{},
				slice2: []transfers.TransactionJson{},
			},
			want: true,
		},
		{
			name: "case 3",
			args: args{
				slice1: []transfers.TransactionJson{
					{
						Id:              1,
						AccountID:       1,
						Amount:          1234.56,
						Bank:            "VCB",
						TransactionType: "deposit",
						CreatedAt:       timeTmp.String(),
					},
				},
				slice2: []transfers.TransactionJson{
					{
						Id:              1,
						AccountID:       1,
						Amount:          1234.56,
						Bank:            "VCB",
						TransactionType: "deposit",
						CreatedAt:       timeTmp.String(),
					},
				},
			},
			want: true,
		},
		{
			name: "case 4",
			args: args{
				slice1: []transfers.TransactionJson{
					{
						Id:              2,
						AccountID:       2,
						Amount:          1234.56,
						Bank:            "VCB",
						TransactionType: "deposit",
						CreatedAt:       timeTmp.String(),
					},
				},
				slice2: []transfers.TransactionJson{
					{
						Id:              1,
						AccountID:       1,
						Amount:          1234.56,
						Bank:            "VCB",
						TransactionType: "deposit",
						CreatedAt:       timeTmp.String(),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTransactionJsonSlice(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("CompareTransactionJsonSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareTransactionJson(t *testing.T) {
	type args struct {
		trans1 *transfers.TransactionJson
		trans2 *transfers.TransactionJson
	}
	timeTmp := time.Now()
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				trans1: nil,
				trans2: nil,
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				trans1: nil,
				trans2: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
			},
			want: false,
		},
		{
			name: "case 3",
			args: args{
				trans1: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
				trans2: nil,
			},
			want: false,
		},
		{
			name: "case 4",
			args: args{
				trans1: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
				trans2: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1111.11,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
			},
			want: false,
		},
		{
			name: "case 5",
			args: args{
				trans1: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
				trans2: &transfers.TransactionJson{
					Id:              20,
					AccountID:       1,
					Amount:          1234.56,
					Bank:            "VCB",
					TransactionType: "deposit",
					CreatedAt:       timeTmp.String(),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTransactionJson(tt.args.trans1, tt.args.trans2); got != tt.want {
				t.Errorf("CompareTransactionJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
