package main

import (


	"github.com/HosseinForouzan/url-shortening-service/handler"
	"github.com/HosseinForouzan/url-shortening-service/repository/psql"
	"github.com/HosseinForouzan/url-shortening-service/repository/redis"
	"github.com/HosseinForouzan/url-shortening-service/shorten"
)

func main() {


	


	conn := psql.New()
    rdb := redis.New()
	shortenSvc := shorten.New(conn, rdb)

	shortHandler := handler.New(shortenSvc)
	shortHandler.SetRoutes()



	
}

