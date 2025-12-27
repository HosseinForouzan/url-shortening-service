package main

import (
	"context"
	"fmt"
	"time"

	"github.com/HosseinForouzan/url-shortening-service/handler"
	"github.com/HosseinForouzan/url-shortening-service/repository/psql"
	"github.com/HosseinForouzan/url-shortening-service/repository/redis"
	"github.com/HosseinForouzan/url-shortening-service/shorten"
)

func main() {

    rdb := redis.New()
    res, err := rdb.RedisSet(context.Background(), "hossein", "salam", 1 * time.Hour)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(res)

    a, _ := rdb.RedisGet(context.Background(), "hossein")
    fmt.Println(a)

	


	conn := psql.New()
	shortenSvc := shorten.New(conn)

	shortHandler := handler.New(shortenSvc)
	shortHandler.SetRoutes()



	
}

