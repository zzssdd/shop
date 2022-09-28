package controller

import (
	"context"
	"shop_payment/model"
	"shop_payment/proto/payment"
)

type Payment struct {
}

func (p *Payment) PaymentList(ctx context.Context, in *payment.ListRequest, out *payment.ListResponse) error {
	email := in.Email
	pageSize := int(in.PageSize)
	currentPage := int(in.CurrentPage)
	list, count, err := model.PaymentList(email, pageSize, currentPage)
	if err != nil {
		out.Code = 500
		out.Msg = "获取订单失败"
		return err
	}
	rep_payments := []*payment.PaymentInfo{}
	for _, v := range list {
		rep_payment := &payment.PaymentInfo{
			Id:       int32(v.ID),
			Pid:      int32(v.Pid),
			Name:     v.Pro.Name,
			BuyPrice: v.BuyPrice,
			Pic:      v.Pro.Pic,
			AddTime:  v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_payments = append(rep_payments, rep_payment)
	}
	out.Code = 200
	out.Msg = "获取订单成功"
	out.Payments = rep_payments
	out.Total = int32(count)
	return nil
}
func (p *Payment) PaymentAdd(ctx context.Context, in *payment.AddRequest, out *payment.Response) error {
	pid := int(in.Pid)
	email := in.Email
	price := in.BuyPrice
	data := &model.Payment{
		Pid:      pid,
		BuyPrice: price,
		Email:    email,
	}
	err := model.PaymentAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "添加订单失败"
		return err
	}
	out.Code = 200
	out.Msg = "添加订单成功"
	return nil
}
func (p *Payment) PaymentDel(ctx context.Context, in *payment.DelRequest, out *payment.Response) error {
	id := int(in.Id)
	err := model.PaymentDel(id)
	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return err
	}
	out.Code = 200
	out.Msg = "退货成功"
	return nil
}
