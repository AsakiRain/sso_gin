package db

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var CACHE *cache.Cache

func SetupGoCache() {
	c := cache.New(10*time.Minute, 10*time.Minute)
	CACHE = c
}
