package main

import (
	"fmt"
	"io"
	"net/http"
)

// func hello(w http.ResponseWriter, req *http.Request) {
// 	ctx := req.Context()
// 	fmt.Println("server: hello handler started")
// 	defer fmt.Println("server: hello handler ended")

// 	select {
// 	case <-time.After(10 * time.Second):
// 		fmt.Fprintf(w, "hello\n")
// 	case <-ctx.Done():

// 		err := ctx.Err()
// 		fmt.Println("server:", err)
// 		internalError := http.StatusInternalServerError
// 		http.Error(w, err.Error(), internalError)
// 	}

// }

// func mainpage(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Response: this is main page\n")
// }

// // func headers(w http.ResponseWriter, req *http.Request) {
// // 	for name, headers := range req.Header {
// // 		for _, h := range headers {
// // 			fmt.Fprintf(w, "%v: %v\n", name, h)
// // 		}
// // 	}
// // }

// func main() {
// 	http.HandleFunc("/hello", hello)
// 	http.HandleFunc("/", mainpage)
// 	// http.HandleFunc("/headers", headers)
// 	http.ListenAndServe(":8090", nil)
// }
// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// // create response binary data
	// data := []byte("Hello World!") // slice of bytes
	// // write `data` to response
	// res.Write(data)
	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "My")
	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " Precious! ")
	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func main() {
	// create a new handler
	handler := HttpHandler{}
	// listen and serve
	http.ListenAndServe(":9000", handler)
}
