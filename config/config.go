package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var Conf Config

type Config struct {
	Server   server   `json:"server"`
	Database database `json:"database"`
}
type server struct {
	Port string `json:"port"`
}
type database struct {
	Mysql mysql `json:"mysql"`
	Redis redis `json:"redis"`
}
type mysql struct {
	Url string `json:"url"`
}
type redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func init() {
	log.Print("读取配置文件.....")
	stream, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Printf("打开文件出错：%v\n", err)
	}
	Conf = Config{}
	err2 := json.Unmarshal(stream, &Conf)
	if err2 != nil {
		log.Fatal("配置文件解析出错:", err2)
	}
}
