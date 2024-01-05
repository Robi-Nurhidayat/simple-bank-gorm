package handler

import (
	"net/http"

	"github.com/Robi-Nurhidayat/simple-bank-gorm/account"
	"github.com/Robi-Nurhidayat/simple-bank-gorm/utils"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	service account.AccountService
}

func NewAccountHandler(service account.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

func (a *AccountHandler) CreateNewAccount(c *gin.Context) {

	var request account.CreateAccountRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	newAccount, err := a.service.CreateAccount(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := account.AccountFormatter(newAccount)

	c.JSON(http.StatusCreated, utils.ApiResponse("success", http.StatusOK, response))

}

func (a *AccountHandler) ListAccount(c *gin.Context) {

	accounts, err := a.service.ListAccounts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := account.ListAccountFormatter(accounts)

	c.JSON(http.StatusOK, utils.ApiResponse("success", http.StatusOK, response))
}


func (a *AccountHandler) UpdateAccount(c *gin.Context) {

	var idUri account.IdFromUri

	err := c.ShouldBindUri(&idUri)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	accountFound,err := a.service.GetAccountByID(idUri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	if accountFound.ID == 0 {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	var req account.UpdateAccountRequest

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse("please ini semua data",http.StatusBadRequest,[]account.Account{}))
		return
	}

	accountFound.Owner = req.Owner
	accountFound.Balance = req.Balance
	accountFound.Currency = req.Currency

	updateAccount,err := a.service.UpdateAccount(accountFound)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse("please ini semua data",http.StatusBadRequest,[]account.Account{}))
		return
	}

	c.JSON(http.StatusOK,utils.ApiResponse("successfully updated account",http.StatusOK,account.AccountFormatter(updateAccount)))
}


func (a *AccountHandler) DeleteById(c *gin.Context) {

	var idUri account.IdFromUri

	err := c.ShouldBindUri(&idUri)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	accountFound,err := a.service.GetAccountByID(idUri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	if accountFound.ID == 0 {
		if err != nil {
			c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
			return
		}
	}

	err = a.service.DeleteById(accountFound.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse("failed to detele data",http.StatusBadRequest,[]account.Account{}))
		return
	}

	c.JSON(http.StatusOK,utils.ApiResponse("successfully delete account",http.StatusOK,[]account.Account{}))


}


func (a *AccountHandler) GetAccountByID(c *gin.Context) {

	var idUri account.IdFromUri

	err := c.ShouldBindUri(&idUri)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	accountFound,err := a.service.GetAccountByID(idUri.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
		return
	}

	if accountFound.ID == 0 {
		if err != nil {
			c.JSON(http.StatusNotFound, utils.ApiResponse("not found that id",http.StatusNotFound,[]account.Account{}))
			return
		}
	}

	c.JSON(http.StatusOK,utils.ApiResponse("successfully delete account",http.StatusOK,account.AccountFormatter(accountFound)))


}