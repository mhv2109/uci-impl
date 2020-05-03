package config

import "sync"

// Configuration interface is similar in behavior to that of a map, and, in fact,
// that's exactly what the default interface is: a thread-safe wrapper around a
// map.
type Configuration interface {
	Get(string) *string // get the value of a Configuration parameter, returning nil if not set
	Set(string, string) // set the value of a Configuration parameter
}

type configurationImpl struct {
	config  map[string]*string
	rwMutex *sync.RWMutex
}

// NewConfiguration returns a newly inititalize implementation of Configuration
// interface.
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
