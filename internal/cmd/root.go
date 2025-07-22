/*
Copyright © 2025 Jeyrce Lu <jeyrce@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	initCfgFile string // 启动时配置，更新后无法重新加载
	runCfgFile  string // 运行时配置，可以热加载
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "internal",
	Short: "quick start you application",
	Long:  ``,
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&initCfgFile, "init-config", "ic", "启动时配置")
	rootCmd.PersistentFlags().StringVar(&runCfgFile, "run-config", "rc", "运行时配置(热更新)")
	rootCmd.Flags().BoolP("version", "v", false, "软件版本信息")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if initCfgFile != "" {

	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".internal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".internal")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
