package controller

import (
	"github.com/gin-gonic/gin"
	"jixiangup.me/examples/gin/internal/repository"
	"strconv"
)

type AccountController struct {
	repository *repository.AccountRepository
}

func NewAccountController(repository *repository.AccountRepository) *AccountController {
	return &AccountController{repository: repository}
}

func (c AccountController) GetUserById(context *gin.Context) {
	sourceId := context.Param("id")
	id, err := strconv.ParseInt(sourceId, 10, 64)
	account, err := c.repository.GetUserById(id)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(200, account)
}
