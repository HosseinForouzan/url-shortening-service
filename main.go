package main

import (
	"fmt"

	"github.com/HosseinForouzan/url-shortening-service/repository/psql"
	"github.com/HosseinForouzan/url-shortening-service/shorten"
)

func main() {

	conn := psql.New()
	shortenSvc := shorten.New(conn)
	a, err := shortenSvc.CreateService(shorten.ShortenRequest{URL: "go.dev.pkg"})
	if err != nil {
		fmt.Println(err.Error())
	}


	fmt.Println(a)

}