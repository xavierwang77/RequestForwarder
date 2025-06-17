package config

import (
	"OffMetaCore/cmn/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Init() {
	logger = log.GetLogger()

	err := initViper()
	if err != nil {
		logger.Fatal("[ FAIL ] failed to init viper", zap.Error(err))
	}

	log.MiniLogger.Info("[ OK ] config module initialed", zap.String("path", viper.ConfigFileUsed()))
}

func initViper() error {
	// 读取配置文件
	viper.SetConfigName(".config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../..")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("init config failed", zap.Error(err))
		return err
	}

	return nil
}
