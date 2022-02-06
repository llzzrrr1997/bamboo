package env

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
)

type BambooEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *BambooEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewBambooEnv
}

// Boot will called when the service instantiate
func (provider *BambooEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BambooEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BambooEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *BambooEnvProvider) Name() string {
	return contract.EnvKey
}
