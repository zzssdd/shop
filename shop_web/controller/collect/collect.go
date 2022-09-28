package collect

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_collect/proto/collection"
	"strconv"
)

func CollectList(c *gin.Context) {
	email, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	list_data := &collection.ListRequest{
		Email:       email.(string),
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	}
	rep, _ := CollectServer.CollectionList(context.TODO(), list_data)
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.Lists,
		"count": rep.Count,
	})
}

func CollectAdd(c *gin.Context) {
	email, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "请先登录",
		})
		return
	}
	pid, _ := strconv.Atoi(c.PostForm("pid"))

	add_data := &collection.AddRequest{
		Email: email.(string),
		Pid:   int32(pid),
	}

	rep, _ := CollectServer.CollectionAdd(context.TODO(), add_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func CollectDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	del_data := &collection.DelRequest{
		Id: int32(id),
	}

	rep, _ := CollectServer.CollectionDel(context.TODO(), del_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
