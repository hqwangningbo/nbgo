package global

import (
	"github.com/hqwangningbo/nbgo/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
	Redis  *conf.RedisClient
)
