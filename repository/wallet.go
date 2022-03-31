package repository

import (
	"context"
	"errors"
	"github.com/Toflex/Wallet-API/errs"
	"github.com/Toflex/Wallet-API/models"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

var Wallets = []models.Wallet{
	{
		WalletID: 10,
		Balance: 230.89,
	},
	{
		WalletID: 11,
		Balance: 10.02,
	},
}

func (r *DefaultRepo) GetWalletInformation(walletId int) (*models.Wallet, *errs.AppError) {
	wallet:= &models.Wallet{}
	err := r.DB.First(wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.NotFoundError("Wallet not found")
	} else if err != nil {
		return nil, errs.UnexpectedError("Unable to fetch wallet, please try again!")
	}

	err=r.AddWalletBalanceToCache(strconv.Itoa(wallet.WalletID), wallet.Balance)
	if err != nil {
		log.Errorf("Error adding to cache %s", err.Error())
	}

	return wallet, nil
}

func (r *DefaultRepo) CreditWalletBalance(wallet *models.Wallet, amount decimal.Decimal) error {
	balance := amount.Add(decimal.NewFromFloat(wallet.Balance)).Round(2).InexactFloat64()
	err:=r.DB.Model(&wallet).Update("balance", balance).Error
	if err != nil {
		return err
	}
	err=r.AddWalletBalanceToCache(strconv.Itoa(wallet.WalletID), balance)
	if err != nil {
		log.Errorf("Error adding to cache %s", err.Error())
	}
	return nil
}

func (r *DefaultRepo) DebitWalletBalance(wallet *models.Wallet, amount decimal.Decimal) error {
	balance := decimal.NewFromFloat(wallet.Balance).Sub(amount).Round(2).InexactFloat64()
	err:=r.DB.Model(&wallet).Update("balance", balance).Error
	if err != nil {
		return err
	}

	strconv.Itoa(wallet.WalletID)

	err=r.AddWalletBalanceToCache(strconv.Itoa(wallet.WalletID), balance)
	if err != nil {
		log.Errorf("Error adding to cache %s", err.Error())
	}
	return nil
}

func (r *DefaultRepo) GetWalletInformationFromCache(walletId string) (float64, error) {
	ctx:=context.Background()

	val, err := r.RDB.Get(ctx, walletId).Float64()
	if err == redis.Nil{
		return 0, errors.New("wallet not found in cache")
	}else if err != nil {
		return 0, err
	}

	return val, nil
}


func (r *DefaultRepo) AddWalletBalanceToCache(walletId string, balance float64) error {
	ctx:=context.Background()

	err := r.RDB.Set(ctx, walletId, balance, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

