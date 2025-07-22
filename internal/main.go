/*
Copyright © 2025 Jeyrce Lu <jeyrce@gmail.com>
*/
package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/0x547d/quickit/internal/cmd"
)

func init() {
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
}

func main() {
	cmd.Execute()
}
