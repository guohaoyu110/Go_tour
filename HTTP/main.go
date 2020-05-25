package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com") // cannot directly write google.com
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("resp: ", resp)

	bs := make([]byte, 99999)
	resp.Body.Read(bs) // response body take the byte
	fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}

	io.Copy(lw, resp.Body)
	// what is this copy thing work?
	// the implementation of io.Copy
}
func (logWriter) Write(bs []byte) (int, error) {
	// implementation inside of write
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}

// in go, we can assemble different interfaces to form an interface
// type ReadCloser interface {
//     Reader
//     Closer
// }
