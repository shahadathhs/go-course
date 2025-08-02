package basic

import (
	console "fmt"
	https "net/http"
)

func Import() {
	console.Println("Hello, GO Standard Library")

	res, err := https.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		console.Println("Error", err)
		return
	}
	defer res.Body.Close()

	console.Println("HTTP Response Status", res.Status)
}
