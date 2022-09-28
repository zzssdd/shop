package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"os"
)

type RedisConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var Con redis.Conn

func LoadRedisConf() (Conf *RedisConf) {
	file, err := os.Open("conf/redis_conf.json")
	if err != nil {
		panic(err)
	}
	byte_data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_data, &Conf)
	if err != nil {
		panic(err)
	}
	return
}

func init() {
	conf := LoadRedisConf()
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	Con, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
}

func Get(email string) (is_ok interface{}, err error) {
	is_ok, err = Con.Do("GET", email)
	return
}

func Set(email string, data string) error {
	_, err := Con.Do("SET", email, data)
	return err
}
