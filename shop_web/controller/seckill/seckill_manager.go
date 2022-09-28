package seckill

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_product/proto/seckillManager"
	"strconv"
	"time"
)

func SeckillList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "6"))
	seckill_data := &seckillManager.SeckillsRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	}
	rep, _ := ManagerService.SeckillList(context.TODO(), seckill_data)
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"data":  rep.List,
		"count": rep.Count,
	})
}
func SeckillAdd(c *gin.Context) {
	price, _ := strconv.ParseFloat(c.PostForm("secPrice"), 64)
	num, _ := strconv.Atoi(c.PostForm("num"))
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	start := c.PostForm("startTime")
	end := c.PostForm("endTime")
	desc := c.PostForm("secDesc")
	add_data := &seckillManager.Seckill{
		Pid:       int32(pid),
		Num:       int32(num),
		StartTime: start,
		EndTime:   end,
		SecDesc:   desc,
		SecPrice:  float32(price),
	}
	rep, _ := ManagerService.SecKillAdd(context.TODO(), add_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
	return
}

func SeckillDel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	del_data := &seckillManager.IdRequest{
		Id: int32(id),
	}
	rep, _ := ManagerService.SeckillDel(context.TODO(), del_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func SeckillInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	info_data := &seckillManager.IdRequest{
		Id: int32(id),
	}
	rep, _ := ManagerService.SeckillInfo(context.TODO(), info_data)
	c.JSON(http.StatusOK, gin.H{
		"code":    rep.Code,
		"msg":     rep.Msg,
		"seckill": rep.Info,
	})
}
func SeckillEdit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	price, _ := strconv.ParseFloat(c.PostForm("secPrice"), 64)
	num, _ := strconv.Atoi(c.PostForm("num"))
	desc := c.PostForm("secDesc")
	pid, _ := strconv.Atoi(c.PostForm("pid"))
	start := c.PostForm("startTime")
	end := c.PostForm("endTime")
	file, err := c.FormFile("pic")
	if err != nil {
		edit_data := &seckillManager.SeckillEdit{
			Id:        int32(id),
			SecPrice:  float32(price),
			Num:       int32(num),
			SecDesc:   desc,
			Pid:       int32(pid),
			StartTime: start,
			EndTime:   end,
		}
		rep, _ := ManagerService.SeckillDoEdit(context.TODO(), edit_data)
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
	edit_data := &seckillManager.SeckillEdit{
		Id:        int32(id),
		SecPrice:  float32(price),
		Num:       int32(num),
		SecDesc:   desc,
		Pid:       int32(pid),
		StartTime: start,
		EndTime:   end,
		Pic:       "http://localhost:8081" + file_path,
	}
	rep, _ := ManagerService.SeckillDoEdit(context.TODO(), edit_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
