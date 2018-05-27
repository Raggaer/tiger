package main

import (
	"time"

	cache "github.com/robfig/go-cache"
)

func setCacheStorage(expire, cleanup time.Duration) *cache.Cache {
	c := cache.New(expire, cleanup)
	return c
}
