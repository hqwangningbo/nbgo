package global

import (
	"github.com/hqwangningbo/gogofly/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger *zap.SugaredLogger
	DB     *gorm.DB
	Redis  *conf.RedisClient
)
