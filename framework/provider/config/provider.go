package config

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
	"path/filepath"
)

type BambooConfigProvider struct{}

// Register registe a new function for make a service instance
func (provider *BambooConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewBambooConfig
}

// Boot will called when the service instantiate
func (provider *BambooConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BambooConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BambooConfigProvider) Params(c framework.Container) []interface{} {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []interface{}{c, envFolder, envService.All()}
}

/// Name define the name for this service
func (provider *BambooConfigProvider) Name() string {
	return contract.ConfigKey
}
