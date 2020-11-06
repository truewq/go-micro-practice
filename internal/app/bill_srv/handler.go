package bill_srv

import (
	"context"
	internal "go-micro-example-auth/internal/proto"
	"time"
)
type Handler struct {
}

func (*Handler) GetBills(ctx context.Context,
	in *internal.GetBillsRequest, out *internal.GetBillsResponse) error{
	out.Code = 0
	out.Bills = append(out.Bills, &internal.BillInfo{
		BillId:"bill00001",
		ReceiveAddr:"Chengdu, Sichuan, China",
		BuyTime: time.Now().Unix(),
		Status: 0,
		GoodsId: 10001,
		GoodsNum: 1,
		GoodsName: "Go micro in 21 days.",
		PayTime: time.Now().Unix(),
		PostId: "YT1000202010010008",
		PostTime: time.Now().Unix(),
		ReceiveTime: time.Now().Unix(),
	})

	return nil
}

