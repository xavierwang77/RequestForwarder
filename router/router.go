package router

import (
	"OffMetaCore/cmn/log"
	"OffMetaCore/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitRoutes 初始化路由
func InitRoutes(r *gin.Engine) {
	logger := log.GetLogger()

	// 从配置文件读取目标URL
	target := viper.GetString("target")
	if target == "" {
		logger.Fatal("target is required")
	}

	r.Any("/", handler.ReverseProxyHandler(target))

	log.MiniLogger.Info("[ OK ] router initialized", zap.String("target", target))
}
