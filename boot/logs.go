package boot

import (
	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func init() {
	value := GetEnvInfo("develop")
	var logger *zap.Logger
	if value {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	zap.ReplaceGlobals(logger)
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
