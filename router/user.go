package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hqwangningbo/gogofly/api"
	"net/http"
)

func InitUserRoute() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)

		}
		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.GET("", func(context *gin.Context) {
				context.AbortWithStatusJSON(http.StatusOK, gin.H{
					"data": []map[string]any{
						{"id": 1, "name": "zs"},
						{"id": 2, "name": "ls"},
					},
				})
			})

			rgAuthUser.GET("/:id", func(context *gin.Context) {
				context.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zs",
				})
			})
		}
	})
}
