package controllers

import (
	"github.com/Toflex/Wallet-API/dto"
	"github.com/Toflex/Wallet-API/utility"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)


// Authenticate Generate authentication godoc
// @Summary Generate auth token
// @Schemes
// @Description Returns auth token
// @Tags Auth
// @Accept json
// @Produce json
// @Param   default  body	dto.AuthRequest  true  "Login to account"
// @Success 200 {object} dto.Token
// @Router /auth [POST]
func (c *NewController) Authenticate(ct *gin.Context) {
	apiResponse:=utility.NewResponse()

	response:=&dto.Token{}
	requestData:=&dto.AuthRequest{}
	if err := ct.ShouldBindJSON(requestData); err != nil {
		log.Errorf("Request could not be decoded. %s", err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.ValidationError("Invalid request", err.Error()))
		return
	}

	// Generate JWT for user
	token, Tokenerr := c.Service.GenerateToken(requestData.EmailAddress)
	if Tokenerr != nil {
		log.Error("An error occurred when generating jwt token. %s", Tokenerr.Error())
		ct.JSON(http.StatusInternalServerError, apiResponse.Error("Authentication failed, please try again!"))
		return
	}

	response.Token = token

	ct.JSON(http.StatusOK, apiResponse.Success("Authentication successful", response))
}

// VerifyAuth verify authentication token
// @Summary Verify auth token
// @Description verify auth token
// @Tags Auth
// @Accept json
// @Produce json
// @Param   default  body	dto.Token  true  "Login to account"
// @Success 200
// @Failure 401
// @Router /auth/verify [POST]
func (c *NewController) VerifyAuth(ct *gin.Context) {
	apiResponse:=utility.NewResponse()

	requestData:=&dto.Token{}
	if err := ct.ShouldBindJSON(requestData); err != nil {
		log.Errorf("Request could not be decoded. %s", err.Error())
		ct.JSON(http.StatusBadRequest, apiResponse.ValidationError("Invalid request", err.Error()))
		return
	}

	// Generate JWT for user
	token, Tokenerr := c.Service.ValidateToken(requestData.Token)
	if Tokenerr != nil {
		log.Error("An error occurred when verifying jwt token. %s", Tokenerr.Error())
		ct.Status(http.StatusUnauthorized)
		return
	}
	if token.Valid {
		ct.Status(http.StatusOK)
		return
	}

	ct.Status(http.StatusUnauthorized)
}