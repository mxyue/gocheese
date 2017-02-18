package util

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/patrickmn/go-cache"
	"time"
)

var cacheBox *cache.Cache

func init() {
	cacheBox = cache.New(10*time.Minute, 30*time.Second)
}

func CacheSet(namespace, key, value string) {
	log.Info("CacheSet key: ", key)
	log.Info("CacheSet value: ", value)
	cacheBox.Set(space_key(namespace, key), value, cache.DefaultExpiration)
	code, _ := cacheBox.Get(space_key(namespace, key))
	log.Info("CacheSet get value: ", code)
}

func CacheGet(namespace, key string) (interface{}, bool) {
	return cacheBox.Get(space_key(namespace, key))
}

func space_key(namespace, key string) string {
	log.Info("space_key:  ", key)
	return fmt.Sprintf("%s-%s", namespace, key)
}

func CacheDelete(namespace, key string) {
	cacheBox.Delete(space_key(namespace, key))
}
