package trace

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
)

type BambooTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *BambooTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewBambooTraceService
}

// Boot will called when the service instantiate
func (provider *BambooTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BambooTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BambooTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *BambooTraceProvider) Name() string {
	return contract.TraceKey
}
