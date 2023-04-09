package service

import (
	"errors"
	"github.com/hqwangningbo/nbgo/dao"
	"github.com/hqwangningbo/nbgo/model"
	"github.com/hqwangningbo/nbgo/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (userService *UserService) Login(userDto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	user := userService.Dao.GetUserByNameAndPassword(userDto.Name, userDto.Password)
	if user.ID == 0 {
		errResult = errors.New("Invalid username or password")
	}

	return user, errResult
}
