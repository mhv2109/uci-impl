package config

import "sync"

type Configuration interface {
	Get(string) *string
	Set(string, string)
}

type configurationImpl struct {
	config  map[string]*string
	rwMutex *sync.RWMutex
}

func NewConfiguration() Configuration {
	return &configurationImpl{
		make(map[string]*string),
		&sync.RWMutex{}}
}

func (config *configurationImpl) Get(key string) *string {
	config.rwMutex.RLock()
	defer config.rwMutex.RUnlock()

	val, ok := config.config[key]
	if !ok {
		return nil
	}
	return val
}

func (config *configurationImpl) Set(key, value string) {
	config.rwMutex.Lock()
	defer config.rwMutex.Unlock()

	config.config[key] = &value
}
