package cmd

import (
	"ahutoj/web/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cfgFilePath string
)

func init() {
	cobra.OnInitialize(initConfig)

	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVar(&cfgFilePath, "conf", "./config.yaml", "conf file")

	_ = viper.BindPFlag("conf", pFlags.Lookup("conf"))
}

func initConfig() {
	logger := utils.GetLogInstance()
	if abs, e := filepath.Abs(cfgFilePath); e != nil {
		logger.Infof("application init from: %s", abs)
	}

	err := utils.ConfigInit(cfgFilePath)
	if err != nil {
		logger.Panicf("conf path(%s) error(%v)", cfgFilePath, err)
		return
	}
	utils.LogInit()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Println(err)
		os.Exit(1)
	}
}
