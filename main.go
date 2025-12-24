package main

import (
	"fmt"

	"github.com/HosseinForouzan/url-shortening-service/repository/psql"
)

func main() {

	conn := psql.New()

	fmt.Println(conn)

}