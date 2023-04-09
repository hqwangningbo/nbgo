package dto

import "github.com/hqwangningbo/nbgo/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required,min=4,max=10" message:"用户名不能为空" min_err:"用户名长度最小为4" max_err:"用户名长度最长为10"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string
	Mobile   string `json:"mobile" form:"mobile" binding:"required,min=11,max=11" min_err:"手机号长度为11位"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=10" message:"密码不能为空" min_err:"密码长度最小为6" max_err:"密码长度最长为10"`
}

func (userAddDto *UserAddDTO) ConvertToUserModel(user *model.User) {
	user.Name = userAddDto.Name
	user.RealName = userAddDto.RealName
	user.Mobile = userAddDto.Mobile
	user.Avatar = userAddDto.Avatar
	user.Password = userAddDto.Password
}
