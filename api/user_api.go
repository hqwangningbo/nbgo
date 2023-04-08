package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hqwangningbo/gogofly/service/dto"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
	}
}

// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登陆详细描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} string "登陆成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func (userApi UserApi) Login(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO

	if err := userApi.BuildRequest(BuildRequestOption{Ctx: ctx, DTO: &userLoginDTO}).GetError(); err != nil {
		return
	}

	userApi.OK(ResponseJson{
		Data: userLoginDTO,
	})
}
