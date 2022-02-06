package app

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
)

// BambooAppProvider 提供App的具体实现方法
type BambooAppProvider struct {
	BaseFolder string
}

// Register 注册BambooApp方法
func (h *BambooAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewBambooApp
}

// Boot 启动调用
func (h *BambooAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *BambooAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *BambooAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *BambooAppProvider) Name() string {
	return contract.AppKey
}
