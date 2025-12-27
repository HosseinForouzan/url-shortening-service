package redis

import "github.com/redis/go-redis/v9"

type Redis struct {
	rdb *redis.Client
}

func New() *Redis {
	    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "strongpassword", // no password set
        DB:       0,  // use default DB
    })


	return &Redis{rdb: rdb}
}

