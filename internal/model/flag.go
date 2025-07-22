package model

import (
	"github.com/spf13/pflag"
)

/*
*
命令行参数选项定义，以及配置获取key关联
*/

const (
	KeyVersion = "version"
)

var (
	VarVersion = pflag.BoolP(KeyVersion, "v", false, "获取应用版本信息")
)
