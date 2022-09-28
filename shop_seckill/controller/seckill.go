package controller

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"shop_seckill/model"
	"shop_seckill/proto"
	"time"
)

type Seckill struct {
}

func (s *Seckill) SeckillAdd(ctx context.Context, in *proto.AddRequest, out *proto.AddResponse) error {
	id := int(in.Id)
	email := in.Email
	seckill, err := model.SeckillInfo(id)
	now := time.Now()
	if err != nil {
		out.Code = 500
		out.Msg = "抢购失败，请稍后"
		return nil
	}
	if now.Before(seckill.StartTime) || now.After(seckill.EndTime) {
		out.Code = 500
		out.Msg = "不在活动时间内"
		return nil
	}
	if seckill.SecNum == 0 {
		out.Code = 500
		out.Msg = "您来晚了，已抢光"
		return nil
	}
	is_ok := model.OrderCheck(email, id)
	if !is_ok {
		out.Code = 500
		out.Msg = "抱歉，每个用户只能抢购一次"
		return nil
	}
	err = model.SeckillSuccess(email, id)
	if err != nil {
		out.Code = 500
		out.Msg = "抢购失败，请重试"
		return nil
	}
	out.Code = 200
	out.Msg = "下单中，请稍后"
	return nil
}

func init() {
	orders, err := model.Ch.Consume("order_queue", "order_consumer", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for order := range orders {
		go OrderApply(order)
	}
}

func OrderApply(order amqp.Delivery) {
	body := string(order.Body)
	var byte_data map[string]interface{}
	err := json.Unmarshal([]byte(body), &byte_data)
	if err != nil {
		panic(err)
	}
	id := int(byte_data["pid"].(float64))
	email := byte_data["email"].(string)
	is_ok := model.OrderCheck(email, id)
	if !is_ok {
		order.Ack(true)
		return
	}
	seckill, err := model.SeckillInfo(id)
	if err != nil {
		order.Ack(true)
		res_map := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败,请重试",
		}
		is_ok, err := model.Get(email)
		if is_ok == "OK" || err != nil {
			return
		}
		byte_data, _ := json.Marshal(res_map)
		model.Set(email, string(byte_data))
		return
	}
	now := time.Now()
	if now.Before(seckill.StartTime) || now.After(seckill.EndTime) {
		order.Ack(true)
		res_map := map[string]interface{}{
			"code": 500,
			"msg":  "不在活动时间内",
		}
		is_ok, err := model.Get(email)
		if is_ok == "OK" || err != nil {
			return
		}
		byte_data, _ := json.Marshal(res_map)
		model.Set(email, string(byte_data))
		return
	}
	if seckill.SecNum == 0 {
		order.Ack(true)
		res_map := map[string]interface{}{
			"code": 500,
			"msg":  "您来晚了，已抢光",
		}
		is_ok, err := model.Get(email)
		if is_ok == "OK" || err != nil {
			return
		}
		byte_data, _ := json.Marshal(res_map)
		model.Set(email, string(byte_data))
		return
	}
	is_ok = model.OrderCheck(email, id)
	if !is_ok {
		order.Ack(true)
		res_map := map[string]interface{}{
			"code": 500,
			"msg":  "每个用户只能抢购一次",
		}
		is_ok, err := model.Get(email)
		if is_ok == "OK" || err != nil {
			return
		}
		byte_data, _ := json.Marshal(res_map)
		model.Set(email, string(byte_data))
		return
	}
	err = model.SeckillSuccess(email, id)
	if err != nil {
		order.Ack(true)
		res_map := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败,请重试",
		}
		is_ok, err := model.Get(email)
		if is_ok == "OK" || err != nil {
			return
		}
		byte_data, _ := json.Marshal(res_map)
		model.Set(email, string(byte_data))
		return
	}
	order.Ack(true)
	res_map := map[string]interface{}{
		"code": 200,
		"msg":  "下单成功",
	}
	res_ok, err := model.Get(email)
	if res_ok == "OK" || err != nil {
		return
	}
	byte_res, _ := json.Marshal(res_map)
	model.Set(email, string(byte_res))
	return
}
