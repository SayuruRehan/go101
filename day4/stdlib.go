package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/posts/1")
	if err != nil {
		fmt.Println("Unable to get from example.com. Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unable to read response body. Error:", err)
		return
	}
	fmt.Println(string(body))
}
