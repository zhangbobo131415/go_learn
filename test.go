package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

//T is a demo
type T struct{ A, B int }

// func main() {
// 	req, _ := http.NewRequest("GET", "https://www.zhihu.com", nil)
// 	cli := http.Client{}
// 	response, _ := cli.Do(req)

// 	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
// 	stdout := os.Stdout
// 	io.Copy(stdout, response.Body)

// 	//返回的状态码
// 	status := response.StatusCode

// 	go fmt.Println(status)

// }

func main() {
	p, _ := filepath.Abs(filepath.Dir(`F:\download\`))
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}
