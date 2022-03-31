package controllers

import (
	"fmt"
	"github.com/Toflex/Wallet-API/dto"
	"github.com/Toflex/Wallet-API/services"
	"github.com/Toflex/Wallet-API/utility"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)


// GetWalletBalance get wallet balance
// @Summary Get wallet balance
// @Description Get wallet balance
// @Tags Wallet
// @Param  wallet_id  path	string false "Wallet ID"
// @Produce json
// @Success 200
// @Router /api/v1/wallets/{wallet_id}/balance [GET]
func (c *NewController) GetWalletBalance(ct *gin.Context)  {
	apiResponse:=utility.NewResponse()
	response:= dto.BalanceResponse{}
	id := ct.Param("wallet_id")

	walletId, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf("Wallet ID is meant to be a valid integer value. %s, %s", id, err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.Error("Incorrect wallet ID"))
		return
	}

	// Fetch Wallet information from cache
	cache, err := c.Repo.GetWalletInformationFromCache(id)
	if err != nil {
		log.Error("Wallet balance does not exist in cache")
		// Get credit wallet info
		wallet, walletErr := c.Repo.GetWalletInformation(walletId)
		if walletErr != nil {
			log.Errorf("Error getting wallet details. %s", walletErr.Message)
			ct.JSON(walletErr.Code, apiResponse.Error(walletErr.Message))
			return
		}

		response.Balance = wallet.Balance
		ct.JSON(http.StatusOK, response)
		return
	}

	response.Balance = cache
	ct.JSON(http.StatusOK, response)
}

// CreditWalletAccount Credit wallet account
// @Summary Credit wallet account
// @Description Credit wallet account
// @Tags Wallet
// @Accept json
// @Produce json
// @Param  wallet_id  path	string false "Wallet ID"
// @Param   default  body	dto.WalletRequest  true  "Login to account"
// @Success 200 {object} utility.Response
// @Failure 401
// @Router /api/v1/wallets/{wallet_id}/credit [POST]
func (c *NewController) CreditWalletAccount(ct *gin.Context)  {
	apiResponse:=utility.NewResponse()
	walletID := ct.Param("wallet_id")

	creditWalletID, err := strconv.Atoi(walletID)
	if err != nil {
		log.Errorf("Wallet ID is meant to be a valid integer value. %s, %s", walletID, err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.Error("Incorrect wallet ID"))
		return
	}

	requestData:=&dto.WalletRequest{}
	if err := ct.ShouldBindJSON(requestData); err != nil {
		log.Errorf("Request could not be decoded. %s", err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.ValidationError("Invalid request", err.Error()))
		return
	}

	amount:= decimal.NewFromFloat(requestData.Amount)

	// Amount is not less than 0
	if services.AmountNotNegative(amount) == false {
		log.Errorf("Credit amount is less than 0. %f", requestData.Amount)
		ct.JSON(http.StatusNotAcceptable, apiResponse.Error("Amount cannot be negative"))
		return
	}

	// Get credit wallet info
	creditWallet, walletErr := c.Repo.GetWalletInformation(creditWalletID)
	if walletErr != nil {
		log.Errorf("Error getting wallet details. %s", walletErr.Message)
		ct.JSON(walletErr.Code, apiResponse.Error(walletErr.Message))
		return
	}

	// Credit wallet
	err = c.Repo.CreditWalletBalance(creditWallet, amount)
	if err != nil {
		log.Errorf("Wallet could not be credited %s", err.Error())
		ct.JSON(http.StatusFailedDependency, apiResponse.Error("Wallet could not be credited at the moment, please try again!"))
		return
	}


	ct.JSON(http.StatusOK, apiResponse.PlainSuccess(fmt.Sprintf("Wallet %s has been credited", walletID)))
}

// DebitWalletAccount Debit wallet account
// @Summary Debit wallet account
// @Description Debit wallet account
// @Tags Wallet
// @Accept json
// @Produce json
// @Param  wallet_id  path	string false "Wallet ID"
// @Param   default  body	dto.WalletRequest  true  "Request param"
// @Success 200 {object} utility.Response
// @Failure 401
// @Router /api/v1/wallets/{wallet_id}/debit [POST]
func (c *NewController) DebitWalletAccount(ct *gin.Context)  {
	apiResponse:=utility.NewResponse()
	walletID := ct.Param("wallet_id")

	debitWalletID, err := strconv.Atoi(walletID)
	if err != nil {
		log.Errorf("Wallet ID is meant to be a valid integer value. %s, %s", walletID, err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.Error("Incorrect wallet ID"))
		return
	}

	requestData:=&dto.WalletRequest{}
	if err := ct.ShouldBindJSON(requestData); err != nil {
		log.Errorf("Request could not be decoded. %s", err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.ValidationError("Invalid request", err.Error()))
		return
	}

	amount := decimal.NewFromFloat(requestData.Amount)

	// Amount is not less than 0
	if services.AmountNotNegative(amount) == false {
		log.Errorf("Debit amount is less than 0. %.2f", requestData.Amount)
		ct.JSON(http.StatusNotAcceptable, apiResponse.Error("Amount cannot be negative"))
		return
	}

	// Get debit wallet info
	debitWallet, walletErr := c.Repo.GetWalletInformation(debitWalletID)
	if walletErr != nil {
		log.Errorf("Error getting wallet details. %s", walletErr.Message)
		ct.JSON(walletErr.Code, apiResponse.Error(walletErr.Message))
		return
	}

	// Does debit wallet have enough credit
	if services.CanDebitWallet(debitWallet.Balance, amount) == false {
		log.Errorf("Insufficient balance, balance cannot be less than zero (0).")
		ct.JSON(http.StatusNotAcceptable, apiResponse.Error("Insufficient balance"))
		return
	}

	// Debit wallet
	err = c.Repo.DebitWalletBalance(debitWallet, amount)
	if err != nil {
		log.Errorf("An error occured when debiting wallet. %s", err.Error())
		ct.JSON(http.StatusFailedDependency, apiResponse.Error("Wallet could not be debited at the moment, please try again!"))
		return
	}

	ct.JSON(http.StatusOK, apiResponse.PlainSuccess(fmt.Sprintf("Wallet %s has been debited", walletID)))
}