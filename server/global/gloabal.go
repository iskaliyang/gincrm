package global

import (
	"gincrm/config"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	REDIS  *redis.Client
	Alipay *alipay.Client
)
