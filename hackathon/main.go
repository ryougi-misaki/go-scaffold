package main

import (
	"fmt"
	"hackathon/config"
	"hackathon/dao/mysql"
	"hackathon/dao/redis"
	"hackathon/routes"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	err = mysql.Init()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	err = redis.Init()
	if err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	routes.Init()
}
