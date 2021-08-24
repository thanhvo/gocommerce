package main

import (
	"fmt"
	"log"
	"user-service/route"
)

func main() {
	fmt.Println("Main Application Starts")
	log.Fatal(route.RunAPI(":8090"))

}
