package controller

import (
	"context"
	"shop_product/model"
	"shop_product/proto/seckillManager"
	"time"
)

type Seckill struct {
}

func (s *Seckill) SeckillList(ctx context.Context, in *seckillManager.SeckillsRequest, out *seckillManager.SeckillsResponse) error {
	currentPage := in.CurrentPage
	pageSize := in.PageSize
	result, count, err := model.SeckillList(int(currentPage), int(pageSize))
	if err != nil {
		out.Code = 500
		out.Msg = "获取秒杀列表成功"
		return nil
	}
	rep_seckills := []*seckillManager.SeckillProduct{}
	for _, cur := range result {
		rep_seckill := &seckillManager.SeckillProduct{
			Id:        int32(cur.Id),
			Name:      cur.Pro.Name,
			Price:     cur.Pro.Price,
			Num:       int32(cur.SecNum),
			SecPrice:  cur.SecPrice,
			StartTime: cur.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   cur.EndTime.Format("2006-01-02 15:04:05"),
			Pic:       cur.Pro.Pic,
			SecDesc:   cur.SecDesc,
			Pid:       int32(cur.Pid),
		}
		rep_seckills = append(rep_seckills, rep_seckill)
	}
	out.List = rep_seckills
	out.Count = int32(count)
	out.Code = 200
	out.Msg = "成功"
	return nil
}

func (s *Seckill) SecKillAdd(ctx context.Context, in *seckillManager.Seckill, out *seckillManager.MsgResponse) error {
	start, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	end, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	data := &model.Seckill{
		Pid:       int(in.Pid),
		SecPrice:  in.SecPrice,
		SecNum:    int(in.Num),
		SecDesc:   in.SecDesc,
		StartTime: start,
		EndTime:   end,
	}
	err := model.SeckillAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "活动添加失败"
		return nil
	}
	out.Code = 200
	out.Msg = "活动添加成功"
	return nil
}
func (s *Seckill) SeckillDel(ctx context.Context, in *seckillManager.IdRequest, out *seckillManager.MsgResponse) error {
	id := int(in.Id)
	err := model.SeckillDel(id)
	if err != nil {
		out.Code = 500
		out.Msg = "活动删除失败"
		return nil
	}
	out.Code = 200
	out.Msg = "活动删除成功"
	return nil
}
func (s *Seckill) SeckillDoEdit(ctx context.Context, in *seckillManager.SeckillEdit, out *seckillManager.MsgResponse) error {
	start, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	end, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	data := &model.Seckill{
		Id:        int(in.Id),
		SecPrice:  in.SecPrice,
		SecNum:    int(in.Num),
		Pid:       int(in.Pid),
		SecDesc:   in.SecDesc,
		StartTime: start,
		EndTime:   end,
	}
	err := model.SeckillEdit(data)
	if err != nil {
		out.Code = 500
		out.Msg = "秒杀商品更新失败"
		return nil
	}
	out.Code = 200
	out.Msg = "秒杀商品更新成功"
	return nil
}
func (s *Seckill) SeckillInfo(ctx context.Context, in *seckillManager.IdRequest, out *seckillManager.SeckillResponse) error {
	id := int(in.Id)
	result, err := model.SeckillInfo(id)
	if err != nil {
		out.Code = 500
		out.Msg = "秒杀商品获取失败"
		return nil
	}
	rep_seckill := &seckillManager.SeckillProduct{
		Id:        int32(result.Id),
		Name:      result.Pro.Name,
		Price:     result.Pro.Price,
		Num:       int32(result.SecNum),
		SecPrice:  result.SecPrice,
		StartTime: result.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:   result.EndTime.Format("2006-01-02 15:04:05"),
		Pic:       result.Pro.Pic,
		SecDesc:   result.SecDesc,
	}
	out.Code = 200
	out.Msg = "秒杀商品获取成功"
	out.Info = rep_seckill
	return nil
}
