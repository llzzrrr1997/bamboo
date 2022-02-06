package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务和获取服务的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务，
	Make(key string) (interface{}, error)
	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) interface{}
	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []interface{}) (interface{}, error)
}

// BambooContainer 是服务容器的具体实现
type BambooContainer struct {
	Container
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instance 存储具体的实例，key为字符串凭证
	instances map[string]interface{}
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

// NewBambooContainer 创建一个服务容器
func NewBambooContainer() *BambooContainer {
	return &BambooContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (bamboo *BambooContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range bamboo.providers {
		name := provider.Name()

		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做了绑定
func (bamboo *BambooContainer) Bind(provider ServiceProvider) error {
	bamboo.lock.Lock()
	key := provider.Name()

	bamboo.providers[key] = provider
	bamboo.lock.Unlock()

	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(bamboo); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(bamboo)
		method := provider.Register(bamboo)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		bamboo.instances[key] = instance
	}
	return nil
}

func (bamboo *BambooContainer) IsBind(key string) bool {
	return bamboo.findServiceProvider(key) != nil
}

func (bamboo *BambooContainer) findServiceProvider(key string) ServiceProvider {
	bamboo.lock.RLock()
	defer bamboo.lock.RUnlock()
	if sp, ok := bamboo.providers[key]; ok {
		return sp
	}
	return nil
}

func (bamboo *BambooContainer) Make(key string) (interface{}, error) {
	return bamboo.make(key, nil, false)
}

func (bamboo *BambooContainer) MustMake(key string) interface{} {
	serv, err := bamboo.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (bamboo *BambooContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return bamboo.make(key, params, true)
}

func (bamboo *BambooContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(bamboo); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(bamboo)
	}
	method := sp.Register(bamboo)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// 真正的实例化一个服务
func (bamboo *BambooContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	bamboo.lock.RLock()
	defer bamboo.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := bamboo.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return bamboo.newInstance(sp, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := bamboo.instances[key]; ok {
		return ins, nil
	}

	// 容器中还未实例化，则进行一次实例化
	inst, err := bamboo.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	bamboo.instances[key] = inst
	return inst, nil
}
