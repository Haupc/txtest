package main

import (
	"txtest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	transferController := controller.GetTransferController()

	r := gin.Default()
	transferRoute := r.Group("/transfer")
	{
		transferRoute.POST("/single", transferController.SingleTransfer)
	}
	r.Run(":8080")
}
