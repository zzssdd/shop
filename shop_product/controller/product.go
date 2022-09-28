package controller

import (
	"context"
	"shop_product/model"
	"shop_product/proto/product"
)

type Product struct {
}

func (p *Product) ProductList(ctx context.Context, in *product.ProductsRequest, out *product.ProductsResponse) error {
	currentPage := int(in.CurrentPage)
	pageSize := int(in.PageSize)
	result, count, err := model.ProductList(currentPage, pageSize)
	if err != nil {
		out.Code = 500
		out.Msg = "商品列表获取失败"
		return nil
	}
	rep_products := []*product.ProductInfo{}
	for _, cur := range result {
		rep_product := &product.ProductInfo{
			Id:         int32(cur.ID),
			Name:       cur.Name,
			Price:      cur.Price,
			Num:        int32(cur.Num),
			Unit:       cur.Unit,
			Pic:        cur.Pic,
			Desc:       cur.Desc,
			CreateTime: cur.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_products = append(rep_products, rep_product)
	}
	out.Code = 200
	out.Msg = "商品列表获取成功"
	out.Products = rep_products
	out.Count = int32(count)
	return nil
}
func (p *Product) ProductAdd(ctx context.Context, in *product.ProductRequest, out *product.MsgResponse) error {
	data := &model.Product{
		Name:  in.Name,
		Price: in.Price,
		Num:   int(in.Num),
		Unit:  in.Unit,
		Pic:   in.Pic,
		Desc:  in.Desc,
	}
	err := model.ProductAdd(data)
	if err != nil {
		out.Code = 500
		out.Msg = "商品增加失败"
		return nil
	}
	out.Code = 200
	out.Msg = "商品增加成功"
	return nil
}
func (p *Product) ProductDel(ctx context.Context, in *product.IdRequest, out *product.MsgResponse) error {
	id := int(in.Id)
	err := model.ProductDel(id)
	if err != nil {
		out.Code = 500
		out.Msg = "商品删除失败"
		return nil
	}
	out.Code = 200
	out.Msg = "商品删除成功"
	return nil
}
func (p *Product) ProductToEdit(ctx context.Context, in *product.IdRequest, out *product.ProductEditResponse) error {
	id := int(in.Id)
	result, err := model.ProductInfo(id)
	if err != nil {
		out.Code = 500
		out.Msg = "商品信息获取失败"
		return nil
	}
	rep_product := &product.ProductInfo{
		Id:         int32(result.ID),
		Name:       result.Name,
		Price:      result.Price,
		Num:        int32(result.Num),
		Unit:       result.Unit,
		Pic:        result.Pic,
		Desc:       result.Desc,
		CreateTime: result.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	out.Code = 200
	out.Msg = "商品信息获取成功"
	out.Info = rep_product
	return nil
}
func (p *Product) ProductDoEdit(ctx context.Context, in *product.EditRequest, out *product.MsgResponse) error {
	id := int(in.Id)
	data := &model.Product{
		Name:  in.Name,
		Price: in.Price,
		Num:   int(in.Num),
		Unit:  in.Unit,
		Pic:   in.Pic,
		Desc:  in.Desc,
	}
	err := model.ProductEdit(id, data)
	if err != nil {
		out.Code = 500
		out.Msg = "商品更新成功"
		return nil
	}
	out.Code = 200
	out.Msg = "商品更新成功"
	return nil
}
