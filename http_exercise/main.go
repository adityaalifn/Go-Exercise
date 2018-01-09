package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}