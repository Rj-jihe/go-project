package main

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	writer.Write([]byte("hello world"))
}
func main() {

	http.HandleFunc("/", Index)
	fmt.Println("web.server listen addr: http://127.0.0.1:80")
	http.ListenAndServe("127.0.0.1:80", nil)

}
