package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/hqwangningbo/nbgo/docs"
	"github.com/hqwangningbo/nbgo/global"
	"github.com/hqwangningbo/nbgo/middleware"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/context"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// 定一个切片，存放想要注册的路由
var (
	gfnRoutes []IFnRegisterRoute
)

// RegisterRoute 将想要注册对路由放到上面的切片中
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// InitRouter 初始化路由系统
func InitRouter() {

	// 创建监听 ctrl + c,应用推出的上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 初始化gin框架，并注册相关路由
	r := gin.Default()
	// 跨域处理
	r.Use(middleware.Cors())
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	//注册自定义校验器
	registerCustValidator()

	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 获取启动端口
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		global.Logger.Info(fmt.Sprintf("Start server port: %s", stPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start server error: %s", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop server error: %s", err.Error()))
		return
	}

	global.Logger.Info(fmt.Sprintf("Stop server success"))
}

func InitBasePlatformRoutes() {
	InitUserRoute()
}

// 注册自定义验证器方法
func registerCustValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				if 0 == strings.Index(value, "a") {
					return true
				}
			}
			return false
		})
	}
}
