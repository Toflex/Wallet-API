package repository

import (
	"github.com/Toflex/Wallet-API/errs"
	"github.com/Toflex/Wallet-API/models"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)


type Repository interface {
	GetWalletInformation(int) (*models.Wallet, *errs.AppError)
	CreditWalletBalance(*models.Wallet, decimal.Decimal) error
	DebitWalletBalance(*models.Wallet, decimal.Decimal) error
	AddWalletBalanceToCache(walletId string, balance float64) error
	GetWalletInformationFromCache(walletId string) (float64, error)
}

type DefaultRepo struct {
	DB *gorm.DB
	RDB *redis.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client) Repository {
	return &DefaultRepo{
		DB: db,
		RDB: rdb,
	}
}
