package cmd

import (
	"OffMetaCore/cmn/config"
	"OffMetaCore/cmn/log"
	"OffMetaCore/handler"
	"OffMetaCore/router"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start all services",
	Long:  `The serve command starts all services of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch debug {
		case true:
			// 设置 Gin 模式为 Debug
			gin.SetMode(gin.DebugMode)
		case false:
			// 设置 Gin 模式为 Release
			gin.SetMode(gin.ReleaseMode)
		default:
			// 设置 Gin 模式为 Release
			gin.SetMode(gin.ReleaseMode)
		}

		// 全局唯一的 Gin 实例
		r := gin.Default()

		r.Use(gin.Logger())
		r.Use(gin.Recovery())

		log.Init(debug)
		logger := log.GetLogger()

		config.Init()
		handler.Init()

		// 引入模块化路由
		router.InitRoutes(r)

		// 读取运行配置
		host := viper.GetString("server.host")
		port := viper.GetString("server.port")
		if host == "" || port == "" {
			logger.Fatal("You need to set server.host and server.port")
		}

		// 启动服务
		err := r.Run(host + ":" + port)
		if err != nil {
			logger.Error("gin run failed", zap.Error(err))
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
