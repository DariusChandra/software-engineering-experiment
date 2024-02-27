package main

import (
	"fmt"
	experiment_coverage "github.com/DariusChandra/software-engineering-experiment/golang/experiment-coverage"
	"github.com/DariusChandra/software-engineering-experiment/golang/go-cache/library"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("codacy", experiment_coverage.Size(0), cache.DefaultExpiration)
	c.Set("codacy2", library.HelloWorld(1), cache.DefaultExpiration)
	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	res, exist := c.Get("codacy")
	if exist {
		fmt.Println(res)
	}
}
