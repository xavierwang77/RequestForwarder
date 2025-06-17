/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "TopnodTenant",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.TopnodTenant.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&phone, "phone", "", "手机号")
	rootCmd.PersistentFlags().StringVar(&msg, "msg", "", "消息内容")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "", "模式")

	rootCmd.PersistentFlags().StringVar(&name, "name", "", "姓名")
	rootCmd.PersistentFlags().StringVar(&gender, "gender", "", "性别")
	rootCmd.PersistentFlags().StringVar(&birth, "birth", "", "生日")

	rootCmd.PersistentFlags().StringVar(&id, "id", "", "主体ID")
	rootCmd.PersistentFlags().StringVar(&file, "file", "", "文件名")
	rootCmd.PersistentFlags().Float64Var(&minScore, "minScore", 0, "最低积分")

	rootCmd.PersistentFlags().BoolVar(&enrollInfo, "enrollInfo", false, "包含登记信息")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "启用调试模式")

	rootCmd.PersistentFlags().StringVar(&archive, "archive", "", "藏品名")
	rootCmd.PersistentFlags().StringVar(&criteria, "criteria", "", "排名依据")

	rootCmd.PersistentFlags().StringVar(&tenant, "tenant", "", "租户名")
}
