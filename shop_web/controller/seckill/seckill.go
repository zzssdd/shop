package seckill

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_web/rabbitmq"
	"shop_web/redis"
	"strconv"
)

func Seckill(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	email, exist := c.Get("user")
	defer rabbitmq.Close()
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token不存在",
		})
		return
	}
	//email := c.PostForm("email")
	order_map := map[string]interface{}{
		"email": email,
		"pid":   id,
	}
	byte_data, _ := json.Marshal(order_map)
	msg := string(byte_data)
	err := rabbitmq.Publish(msg)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "抢购失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "抢购中，请稍后",
	})
}

func SeckillResult(c *gin.Context) {
	email, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token不存在",
		})
		return
	}
	//email := c.Query("email")
	result, err := redis.Get(email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "暂未获取到抢购结果",
		})
		return
	}
	var rep map[string]interface{}
	json.Unmarshal([]byte(result), &rep)

	c.JSON(http.StatusOK, gin.H{
		"code": rep["code"],
		"msg":  rep["msg"],
	})
	return
}
