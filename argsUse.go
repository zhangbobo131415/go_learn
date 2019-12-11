package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", "172.0.0.1:8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "Jack", "name")

	flag.Parse()
	//    flag.Usage()
	others := flag.Args()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)
	fmt.Println("other:", others)
	http.ListenAndServe()
}
