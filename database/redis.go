package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	. "main/config"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	log.Print("连接redis缓存数据库")
	rdb = redis.NewClient(&redis.Options{
		Addr:     Conf.Database.Redis.Addr,
		Password: Conf.Database.Redis.Password,
		DB:       Conf.Database.Redis.DB,
	})
	ping := rdb.Ping(ctx)
	if ping == nil {
		log.Fatal("连接redis失败")
	}
}
