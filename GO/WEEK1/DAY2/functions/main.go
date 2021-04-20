package main

import (

	// "errors"
	"fmt"
)

func main() {

	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

//creating func
func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	fmt.Println("Server started ", port)
	fmt.Println("retries ", numberOfRetries)
	// return errors.New("Someting went wrong")
	return port, nil
}
