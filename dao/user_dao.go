package dao

import "github.com/hqwangningbo/nbgo/model"

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

func (userDao UserDao) GetUserByNameAndPassword(username, password string) model.User {
	var user model.User
	userDao.Orm.Model(&user).Where("name=? and password=?", username, password).Find(&user)
	return user
}
