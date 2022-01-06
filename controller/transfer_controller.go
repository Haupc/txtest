package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"txtest/dto"
	"txtest/service"
	"txtest/utils"

	"github.com/gin-gonic/gin"
)

var (
	_transferController    *transferController
	transferControllerLock sync.Mutex
)

func GetTransferController() TransferController {
	if _transferController == nil {
		transferControllerLock.Lock()
		defer transferControllerLock.Unlock()
		if _transferController == nil {
			_transferController = &transferController{
				transferService: service.GetTransferService(),
			}
		}
	}
	return _transferController
}

type TransferController interface {
	SingleTransfer(ctx *gin.Context)
}

type transferController struct {
	transferService service.TransferService
}

func (c *transferController) SingleTransfer(ctx *gin.Context) {
	var singleTransferDto dto.SingleTransferDto
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &singleTransferDto)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	signedTx, err := c.transferService.SingleTransfer(ctx, singleTransferDto.ToAddress, singleTransferDto.Amount)
	if err != nil {
		respose := utils.BuildErrorResponse("Transfer error", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respose)
		return
	}

	respose := utils.BuildResponse(true, "success", signedTx.Hash().Hex())
	ctx.JSON(http.StatusOK, respose)
}
