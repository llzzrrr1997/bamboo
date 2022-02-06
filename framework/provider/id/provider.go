package id

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"github.com/llzzrrr1997/bamboo/framework/contract"
)

type BambooIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *BambooIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewBambooIDService
}

// Boot will called when the service instantiate
func (provider *BambooIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *BambooIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *BambooIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *BambooIDProvider) Name() string {
	return contract.IDKey
}
