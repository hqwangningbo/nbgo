package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hqwangningbo/nbgo/service"
	"github.com/hqwangningbo/nbgo/service/dto"
	"github.com/hqwangningbo/nbgo/utils"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
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

	user, err := userApi.Service.Login(userLoginDTO)
	if err != nil {
		userApi.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.Name)

	userApi.OK(ResponseJson{
		Data: gin.H{
			"user_info": user,
			"token":     token,
		},
	})
}
