package redis

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/soveran/redisurl"
)

// AddToSortedSet inserts a value into the sorted set for the given key at the
// position of the score
func AddToSortedSet(key string, score int64, value string) {
	connection, err := redisurl.ConnectToURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	connection.Do("ZADD", key, score, value)
}

// FetchReversedCollectionFromSortedSet returns count number of values for the
// sorted set of the given key in revere scorew order.
func FetchReversedCollectionFromSortedSet(key string, count int) ([]interface{}, error) {
	connection, err := redisurl.ConnectToURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	return redis.Values(connection.Do("ZREVRANGE", key, 0, count))
}
