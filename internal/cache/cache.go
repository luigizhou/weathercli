package cache

import (
	"sync"
)

var cache = make(map[string]string)
var mutex = sync.RWMutex{}

func GetValueFromCache(city string) (string, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	val, found := cache[city]
	return val, found
}

func SetValueInCache(city, data string) {
	mutex.Lock()
	defer mutex.Unlock()
	cache[city] = data
}
