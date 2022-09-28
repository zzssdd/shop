package controller

import (
	"context"
	"shop_collect/model"
	"shop_collect/proto/collection"
)

type Collection struct {
}

func (c *Collection) CollectionList(ctx context.Context, in *collection.ListRequest, out *collection.ListsResponse) error {
	email := in.Email
	currentPage := int(in.CurrentPage)
	pageSize := int(in.PageSize)
	result, count, err := model.CollectList(email, currentPage, pageSize)
	if err != nil {
		out.Code = 500
		out.Msg = "获取购物车失败"
		return nil
	}
	rep_collections := []*collection.List{}
	for _, cur := range result {
		rep_collection := &collection.List{
			Id:      int32(cur.ID),
			Pid:     int32(cur.Pid),
			Name:    cur.Pro.Name,
			Price:   cur.Pro.Price,
			Unit:    cur.Pro.Unit,
			Pic:     cur.Pro.Pic,
			Desc:    cur.Pro.Desc,
			AddTime: cur.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_collections = append(rep_collections, rep_collection)
	}
	out.Code = 200
	out.Msg = "获取购物车成功"
	out.Lists = rep_collections
	out.Count = int32(count)
	return nil
}
func (c *Collection) CollectionAdd(ctx context.Context, in *collection.AddRequest, out *collection.MsgResponse) error {
	data := model.Collection{
		Email: in.Email,
		Pid:   int(in.Pid),
	}
	err := model.CollectAdd(&data)
	if err != nil {
		out.Code = 500
		out.Msg = "添加购物车失败"
		return nil
	}
	out.Code = 200
	out.Msg = "添加购物车成功"
	return nil
}
func (c *Collection) CollectionDel(ctx context.Context, in *collection.DelRequest, out *collection.MsgResponse) error {
	id := int(in.Id)
	err := model.CollectDel(id)
	if err != nil {
		out.Code = 500
		out.Msg = "删除失败"
		return nil
	}
	out.Code = 200
	out.Msg = "删除成功"
	return nil
}
