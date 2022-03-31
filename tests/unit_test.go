package tests

import (
	"github.com/Toflex/Wallet-API/services"
	"github.com/shopspring/decimal"
	"testing"
)


func Test_Can_Debit_Wallet_Balance_Greater_Than_Debit_Amount_Return_True(t *testing.T) {
	balance := 12.09
	amount := decimal.NewFromInt(10)
	if services.CanDebitWallet(balance, amount) == false {
		t.Errorf("A wallet balance cannot go below 0.")
	}
}

func Test_Can_Debit_Wallet_Balance_Less_Than_Debit_Amount_Return_False(t *testing.T) {
	balance := 12.09
	amount := decimal.NewFromInt(100)
	if services.CanDebitWallet(balance, amount) == true {
		t.Errorf("A wallet balance cannot go below 0.")
	}
}

func Test_Can_Debit_Wallet_Balance_Equal_Debit_Amount_Return_True(t *testing.T) {
	balance := 12.0
	amount := decimal.NewFromInt(12)
	if services.CanDebitWallet(balance, amount) == false {
		t.Errorf("A wallet balance cannot go below 0.")
	}
}

func Test_Amount_Not_Negative_If_Amount_Less_Than_0_Return_False(t *testing.T) {
	amount := decimal.NewFromInt(-20)
	if services.AmountNotNegative(amount) == true {
		t.Errorf("Amounts sent in the credit and debit operations cannot be negative.")
	}
}

func Test_Amount_Not_Negative_If_Amount_Greater_Than_0_Return_True(t *testing.T) {
	amount := decimal.NewFromInt(10)
	if services.AmountNotNegative(amount) == false {
		t.Errorf("Amounts sent in the credit and debit operations cannot be negative.")
	}
}

func Test_Amount_Not_Negative_If_Amount_Equals_0_Return_True(t *testing.T) {
	amount := decimal.NewFromInt(0)
	if services.AmountNotNegative(amount) == false {
		t.Errorf("Amounts sent in the credit and debit operations cannot be negative.")
	}
}