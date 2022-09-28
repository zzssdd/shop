package payment

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_payment/proto/payment"
	"strconv"
)

func PaymentList(c *gin.Context) {
	email, exit := c.Get("user")
	if !exit {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	list_data := &payment.ListRequest{
		Email:       email.(string),
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	}
	rep, err := PaymentServer.PaymentList(context.TODO(), list_data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Payments,
		"count": rep.Total,
	})
}

func PaymentAdd(c *gin.Context) {
	email, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	add_data := &payment.AddRequest{
		Pid:      int32(pid),
		BuyPrice: float32(price),
		Email:    email.(string),
	}
	rep, _ := PaymentServer.PaymentAdd(context.TODO(), add_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

}

func PaymentDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	del_data := &payment.DelRequest{
		Id: int32(id),
	}

	rep, _ := PaymentServer.PaymentDel(context.TODO(), del_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
