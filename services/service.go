package services

import (
	"github.com/Toflex/Wallet-API/configs"
	"github.com/Toflex/Wallet-API/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/shopspring/decimal"
)

type Service interface {
	GenerateToken(string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}


type DefaultService struct {
	repo repository.Repository
	Config *configs.Config
}

func NewService(repo repository.Repository, config *configs.Config) Service {
	return DefaultService{repo: repo, Config: config}
}


// CanDebitWallet ...
// A wallet balance cannot go below 0.
func CanDebitWallet(balance float64, amount decimal.Decimal) bool {
	return decimal.NewFromFloat(balance).Sub(amount).GreaterThanOrEqual(decimal.NewFromFloat(0.0))
}

// AmountNotNegative ...
// Amounts sent in the credit and debit operations cannot be negative.
func AmountNotNegative(amount decimal.Decimal) bool {
	return amount.GreaterThanOrEqual(decimal.NewFromInt(0))
}