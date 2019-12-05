package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	_ "sync"
	"time"

	my_test "./my_pkg_"
)

func init() {
	fmt.Println(my_test.Sum(2, 6666))

}

func main() {
	//my_test.CurlMy()
	fmt.Println(my_test.Sum(2, 6666))

	read := newAlphaReader("Hello! It's 9am, whee is the sun?")
	io.Copy(os.Stdout, read)
	fmt.Println(fmt.Sprintf("fdsf%d",6))
	// fmt.Println()
	// fmt.Println("输入一个字符串：")
	// reader := bufio.NewReader(read)
	// s1, _ := reader.ReadString('\n')
	// fmt.Println("读到的数据：", s1)

	// conn, err := net.Dial("tcp", "172.21.15.216:8000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// done := make(chan struct{})
	// go func() {
	// 	io.Copy(os.Stdout, conn) // NOTE: ignoring errors
	// 	log.Println("done")
	// 	done <- struct{}{} // signal the main goroutine
	// }()
	// mustCopy(conn, os.Stdin)
	// conn.Close()
	// <-done // wat for background goroutine to finish

	// listener, err := net.Listen("tcp", "172.21.15.216:8000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Print(err) // e.g., connection aborted
	// 		continue
	// 	}
	// 	go handleConn(conn) // handle one connection at a time
	// }
}

// AlphaReader is
type AlphaReader struct {
	// 资源
	src string
	// 当前读取到的位
	cur int
}

// 创建一个实例
func newAlphaReader(src string) *AlphaReader {
	return &AlphaReader{src: src}
}

// 过滤函数
func alpha(r byte) byte {
	if r >= 'A' && r <= 'Z' || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

// Read 方法
func (a *AlphaReader) Read(p []byte) (int, error) {
	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}
	fmt.Println(len(p))

	// x 是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度出，缓冲区不补满
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound && a.cur < len(a.src) {
		// 每次读取一个字节，执行过滤函数
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
			n++
		}
		a.cur++
	}
	// 将处理后得到的 buf 内容复制到 p 中
	copy(p, buf)
	return n, nil
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
