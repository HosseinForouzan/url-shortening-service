package main

import (

	"github.com/HosseinForouzan/url-shortening-service/handler"
	"github.com/HosseinForouzan/url-shortening-service/repository/psql"
	"github.com/HosseinForouzan/url-shortening-service/shorten"
)

func main() {

	conn := psql.New()
	shortenSvc := shorten.New(conn)

	shortHandler := handler.New(shortenSvc)
	shortHandler.SetRoutes()

	
}

