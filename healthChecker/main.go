package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 8080
	portString := os.Getenv("PORT")
	if len(portString) > 0 {
		port, _ = strconv.Atoi(portString)
	}

	res, err := http.Get("http://localhost:" + strconv.Itoa(port) + "/status")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("invalid_status: " + res.Status)
		os.Exit(1)
		return
	}

	os.Exit(0)
}