package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var (
	dir = flag.String("dir", "", "file root direction")
	addr = flag.String("addr", ":11015", "file server port")
	)
func main() {
	if *dir == "" {
		fmt.Println("no root dir")
		return
	}
	rootPath, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.Dir(rootPath)))
	if err = http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}

}