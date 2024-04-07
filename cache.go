package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

type allCache struct {
	listCountries *cache.Cache
}

const (
	defaultExpiration = 5 * time.Minute
	purgeTime         = 10 * time.Minute
)

func newCache() *allCache {
	Cache := cache.New(defaultExpiration, purgeTime)
	return &allCache{
		listCountries: Cache,
	}
}

func (c *allCache) read(id string) (item string, ok bool) {
	country, ok := c.listCountries.Get(id)
	//log.Println(country)
	if ok {
		//log.Println("from cache")
		// res, err := json.Marshal(country.(Country))
		res, err := json.Marshal(country)
		//temp, _ := utf8.DecodeRune(res)
		if err != nil {
			log.Fatal("Error")
		}
		name := string(res[1 : len(res)-1])
		newName := strings.Replace(name, "\\u0026", "&", -1)
		//temp, _ := strconv.Unquote(name)
		//log.Println(newName)
		return newName, true
	}
	return "", false
}

func (c *allCache) update(id string, name string) {
	c.listCountries.Set(id, name, cache.DefaultExpiration)
}
