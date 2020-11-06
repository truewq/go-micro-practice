package bill_api

import (
	"context"
	"github.com/gin-gonic/gin"
	internal "go-micro-example-auth/internal/proto"
	"net/http"
)

type Handler struct{}
var(
	Cli2bill internal.BillService
)
func(*Handler) GetBills(c *gin.Context){
	var req internal.GetBillsRequest
	var res internal.GetBillsResponse
	var err error
	status := http.StatusOK
	defer func() {
		if err != nil{
			res.Code = -1
		}
		c.JSON(status, res)
	}()

	err = c.ShouldBindJSON(&req)
	if err != nil{
		status = http.StatusBadRequest
		return
	}

	getBillsRes, err := Cli2bill.GetBills(context.TODO(), &req)
	if err != nil || getBillsRes == nil{
		status = http.StatusInternalServerError
		return
	}

	res = *getBillsRes
}