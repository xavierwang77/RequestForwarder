package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// 查找配置文件
func findConfigFile(configFileName string, levels int) (string, error) {
	// 从当前工作目录开始查找
	dir, err := os.Getwd()
	if err != nil {
		logger.Error("get current working directory failed", zap.Error(err))
		return "", err
	}

	// 遍历上级目录，查找配置文件
	for i := 0; i <= levels; i++ {
		configPath := filepath.Join(dir, configFileName)
		if _, err := os.Stat(configPath); err == nil {
			// 找到文件，返回文件路径
			return configPath, nil
		}
		// 向上一级目录
		dir = filepath.Dir(dir)
	}

	logger.Sugar().Errorf("config file not found within %d levels", levels)
	return "", fmt.Errorf("config file not found within %d levels", levels)
}

// 加载配置文件
func loadConfig(configFileName string) (*viper.Viper, error) {
	// 查找配置文件（往上查找3级）
	configPath, err := findConfigFile(configFileName, 3)
	if err != nil {
		logger.Error("find config file failed", zap.Error(err))
		return nil, err
	}

	// 使用viper读取配置文件
	v := viper.New()
	v.SetConfigFile(configPath)

	// 尝试读取配置文件
	if err := v.ReadInConfig(); err != nil {
		logger.Error("read config file failed", zap.Error(err))
		return nil, err
	}

	// 配置文件读取成功，返回viper实例
	return v, nil
}
