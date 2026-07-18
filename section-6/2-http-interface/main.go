package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
)

type logWriter struct{}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)
}
