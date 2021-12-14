package global

import (
	"RemindMe/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"RemindMe/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Redis  *redis.Client
	Config config.Server
	VP     *viper.Viper
	//Log    *oplogging.Logger
	Log                *zap.Logger
	Timer              timer.Timer = timer.NewTimerTask()
	ConcurrencyControl             = &singleflight.Group{}

	BlackCache local_cache.Cache
)
