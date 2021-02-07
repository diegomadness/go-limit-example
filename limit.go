package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

const LIMIT = 2

func main() {
	c = cache.New(2*time.Second, 10*time.Second)
	http.HandleFunc("/limit", limit)
	http.ListenAndServe(":3001", nil)
}

func limit(w http.ResponseWriter, req *http.Request) {
	apiToken := req.Header.Get("X-API-TOKEN")

	if apiToken == "" {
		fmt.Fprintf(w, "DENIED, missing header X-API-TOKEN")
		return
	}

	if requestThrottled(apiToken) {
		fmt.Fprintf(w, "DENIED, limit is %d per second", LIMIT)
		return
	}

	fmt.Fprintf(w, "ALLOWED by a Go script!")
}

func requestThrottled(apiToken string) bool {
	cacheKey := strconv.FormatInt(time.Now().Unix(), 10) + apiToken

	requestCount := 0
	value, found := c.Get(cacheKey)
	if found {
		requestCount = value.(int)
	}

	requestCount++
	c.Set(cacheKey, requestCount, cache.DefaultExpiration)

	return requestCount > LIMIT
}
