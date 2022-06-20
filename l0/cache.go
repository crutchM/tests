package l0

import (
	"sync"
	"time"
)

type Item struct {
	Value      interface{}
	Created    time.Time
	Expiration int64
}

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanup           time.Duration
	items             map[string]Item
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanup:           cleanupInterval,
	}
	//есди интеравал >0 то запускаем очистку устаревших элементов
	if cleanupInterval > 0 {
		cache.startGC()
	}
	return &cache
}

//запихиваем в кеш значение, возможна перезапись существующего элемента
func (s *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64
	//определение продолжительности жизни айтема
	if duration == 0 {
		duration = s.defaultExpiration
	}
	// Время жизни кеша
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	s.Lock()

	defer s.Unlock()

	s.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}
