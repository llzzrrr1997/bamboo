package services

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
	"os"
)

// BambooConsoleLog 代表控制台输出
type BambooConsoleLog struct {
	BambooLog
}

// NewBambooConsoleLog 实例化BambooConsoleLog
func NewBambooConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &BambooConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
