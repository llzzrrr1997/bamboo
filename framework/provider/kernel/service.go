package kernel

import (
	"github.com/llzzrrr1997/bamboo/framework/gin"
	"net/http"
)

// 引擎服务
type BambooKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewBambooKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &BambooKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *BambooKernelService) HttpEngine() http.Handler {
	return s.engine
}
