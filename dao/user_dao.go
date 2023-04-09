package dao

import (
	"github.com/hqwangningbo/nbgo/model"
	"github.com/hqwangningbo/nbgo/service/dto"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			NewBaseDao(),
		}
	}
	return userDao
}

func (userDao *UserDao) GetUserByNameAndPassword(username, password string) model.User {
	var user model.User
	userDao.Orm.Model(&user).Where("name=? and password=?", username, password).Find(&user)
	return user
}

func (userDao *UserDao) AddUser(userAddDto *dto.UserAddDTO) error {
	var user model.User
	userAddDto.ConvertToUserModel(&user)

	err := userDao.Orm.Save(&user).Error

	if err == nil {
		userAddDto.ID = user.ID
		userAddDto.Password = ""
	}
	return err
}

func (userDao *UserDao) CheckUserNameExist(username string) bool {
	var totalAmount int64
	userDao.Orm.Model(&model.User{}).Where("name=?", username).Count(&totalAmount)
	return totalAmount > 0
}

func (userDao *UserDao) GetUserById(id uint) (model.User, error) {
	var user model.User
	err := userDao.Orm.First(&user, id).Error
	return user, err
}
