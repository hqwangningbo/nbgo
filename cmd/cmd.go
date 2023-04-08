package cmd

import (
	"fmt"
	"github.com/hqwangningbo/nbgo/conf"
	"github.com/hqwangningbo/nbgo/global"
	"github.com/hqwangningbo/nbgo/router"
	"github.com/hqwangningbo/nbgo/utils"
)

func Start() {
	var initError error
	// 初始化系统配置
	conf.InitConfig()
	//初始化日志组件
	global.Logger = conf.InitLogger()

	//初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initError = utils.AppendError(initError, err)
	}

	//初始化redis
	redis, err := conf.InitRedis()
	global.Redis = redis
	if err != nil {
		initError = utils.AppendError(initError, err)
	}

	_ = global.Redis.Delete("password", "name")
	fmt.Println(global.Redis.Get("password"))

	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}
	}

	//初始化系统路由
	router.InitRouter()
}

func Clean() {
	fmt.Println("----------------Clean")
}
