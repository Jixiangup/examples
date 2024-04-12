package controller

import (
	"github.com/gin-gonic/gin"
	error2 "jixiangup.me/examples/gin/internal/error"
	"jixiangup.me/examples/gin/internal/service"
	"net/http"
	"strconv"
)

type AccountController struct {
	service *service.AccountService
}

func NewAccountController(service *service.AccountService) *AccountController {
	return &AccountController{service: service}
}

func (c AccountController) AccountDetail(context *gin.Context) {
	sourceId := context.Param("id")
	id, err := strconv.ParseInt(sourceId, 10, 64)
	if id != 0 {
		Error(context, error2.NewValidation("Id cannot by empty"))
		return
	}

	account, err := c.service.GetAccountDetail(id)
	if err != nil {
		Error(context, error2.NewUnknownMistake(err.Error()))
		return
	}
	context.JSON(http.StatusOK, account)
}
