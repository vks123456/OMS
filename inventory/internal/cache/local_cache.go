package cache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache"
)

type AppCache struct {
	BigCache *bigcache.BigCache
}

func (a *AppCache) Initialize(config bigcache.Config) error {
	bigCache, err := bigcache.NewBigCache(config)
	a.BigCache = bigCache
	return err
}

func (a *AppCache) InitUpdateCacheScheduler(stopScheduler chan bool) {

	/*
		Update cache here
	*/
	updateCacheScheduler := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-stopScheduler:
				updateCacheScheduler.Stop()
				return
			case <-updateCacheScheduler.C:
				/*
					Update cache here
				*/
			}
		}
	}()
}

func (a *AppCache) Get(key string) ([]byte, error) {
	if a.BigCache == nil {
		return nil, errors.New("Cache layer not properly initialized")
	}
	return a.BigCache.Get(key)
}

func (a *AppCache) Set(key string, value []byte) error {
	if a.BigCache == nil {
		return errors.New("Cache layer not properly initialized")
	}
	return a.BigCache.Set(key, value)
}
