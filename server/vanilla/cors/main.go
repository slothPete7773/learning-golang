package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Hello World\n")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	fmt.Printf("got / request.\n")
	fmt.Printf("Body: \n%s\n", body)
	fmt.Printf("Header: \n%s\n", r.Header)
	for k, v := range r.Header {
		fmt.Printf("k: %v\nv: %v\n---------\n", k, v)
		// for k, v := range m {
		// }
	}
	io.WriteString(w, "This is my website!\n")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server listening on port http://localhost:8081 .")
	http.ListenAndServe(":8081", nil)
}
