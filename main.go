/*
@File: main.go
@Contact: lucien@lucien.ink
@Licence: (C)Copyright 2019 Lucien Shui

@Modify Time      @Author    @Version    @Description
------------      -------    --------    -----------
2021-01-02 17:00  Bryce      1.0         Init
*/
package main

import (
	"PasteMeGoAccount/config"
	"PasteMeGoAccount/flag"
	"PasteMeGoAccount/handler"
	"PasteMeGoAccount/model"
	"PasteMeGoAccount/router"
	"PasteMeGoAccount/vail"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/wonderivan/logger"
)

func init() {
	initDB()
	initRedis()
	vail.InitValidator()
	initRouter()
}

func main() {
	// todo:bzy 参数校验、通用错误处理、缓存持久化、事务处理、服务鉴权
}

func initRouter() {
	// 路由初始化
	serverConfig := config.Get()
	ginRouter := gin.Default()

	// 路由加载
	router.LoadAccountRouter(ginRouter)
	router.LoadPermissionRouter(ginRouter)

	err := ginRouter.Run(fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port))
	if err != nil {
		logger.Painc("Run server failed: " + err.Error())
	}
}

func initDB() {
	// 数据库连接初始化
	dbConfig := config.Get().Database
	gormDB, err := gorm.Open("mysql", connectFormat(dbConfig.Username, dbConfig.Password, "tcp", dbConfig.Server, dbConfig.Port, dbConfig.Database))

	if err != nil {
		logger.Painc("Connect to MySQL failed: " + err.Error())
	} else {
		logger.Info("MySQL connected")
		if flag.Debug {
			logger.Warn("Running in debug mode, database execute will be displayed")
			gormDB = gormDB.Debug()
		}
	}
	gormDB.LogMode(true)

	// 数据表加载
	model.LoadAccountModel(gormDB)
	model.LoadBasicInfoModel(gormDB)
	model.LoadResourceModel(gormDB)
	model.LoadPermissionModel(gormDB)
	model.LoadOpLogModel(gormDB)
}

// 数据库连接参数格式化
func connectFormat(
	username string,
	password string,
	network string,
	server string,
	port uint16,
	database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&loc=Local", username, password, network, server, port, database)
}

func initRedis() {
	// Redis 连接初始化
	redisConfig := config.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr: redisConfig.Address,
		DB: redisConfig.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		logger.Painc("redis 连接失败：", pong, err)
		return
	}
	logger.Info("redis 连接成功：", pong)

	handler.LoadRedisClient(client)
}
