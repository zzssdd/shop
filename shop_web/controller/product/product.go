package product

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_product/proto/product"
	"strconv"
	"time"
)

func ProductList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	list_data := &product.ProductsRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	}
	rep, _ := ProductService.ProductList(context.TODO(), list_data)
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Products,
		"count": rep.Count,
	})
}

func ProductAdd(c *gin.Context) {
	name := c.PostForm("name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	num, _ := strconv.Atoi(c.PostForm("num"))
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")
	file, err := c.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		add_data := &product.ProductRequest{
			Name:  name,
			Price: float32(price),
			Num:   int32(num),
			Unit:  unit,
			Desc:  desc,
		}
		rep, _ := ProductService.ProductAdd(context.TODO(), add_data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	err = c.SaveUploadedFile(file, file_path)
	add_data := &product.ProductRequest{
		Name:  name,
		Price: float32(price),
		Num:   int32(num),
		Unit:  unit,
		Desc:  desc,
		Pic:   "http://localhost:8081/" + file_path,
	}
	rep, _ := ProductService.ProductAdd(context.TODO(), add_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ProductDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Println(id)
	del_data := &product.IdRequest{
		Id: int32(id),
	}

	rep, _ := ProductService.ProductDel(context.TODO(), del_data)

	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func ProductInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	toEdit_data := &product.IdRequest{
		Id: int32(id),
	}

	rep, _ := ProductService.ProductToEdit(context.TODO(), toEdit_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
		"data": rep.Info,
	})
}

func ProductEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	num, _ := strconv.Atoi(c.PostForm("num"))
	unit := c.PostForm("unit")
	desc := c.PostForm("desc")
	file, err := c.FormFile("pic")
	if err != nil {
		edit_data := &product.EditRequest{
			Id:    int32(id),
			Name:  name,
			Price: float32(price),
			Num:   int32(num),
			Unit:  unit,
			Desc:  desc,
		}
		rep, _ := ProductService.ProductDoEdit(context.TODO(), edit_data)
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	unix_int64 := time.Now().Unix()
	unix_str := strconv.FormatInt(unix_int64, 10)
	file_path := "upload/" + unix_str + file.Filename
	c.SaveUploadedFile(file, file_path)
	edit_data := &product.EditRequest{
		Id:    int32(id),
		Name:  name,
		Price: float32(price),
		Num:   int32(num),
		Unit:  unit,
		Desc:  desc,
		Pic:   "http://localhost:8081" + file_path,
	}
	rep, _ := ProductService.ProductDoEdit(context.TODO(), edit_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
