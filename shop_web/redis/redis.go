package redis

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

var (
	Conn redis.Conn
	err  error
)

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
	dsn := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	Conn, err = redis.Dial("tcp", dsn)
}

func Get(email interface{}) (string, error) {
	return redis.String(Conn.Do("get", email))
}

func Close() {
	Conn.Close()
}
