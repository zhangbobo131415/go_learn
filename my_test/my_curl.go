package my_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// CurlMy is test
func CurlMy() {
	// 从Web服务器得到响应
	r, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	//从Body复制到Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
