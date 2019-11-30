package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

type alphaReader struct {
	// 资源
	src string
	// 当前读取到的位置
	cur int
}

// 创建一个实例
func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

// 过滤函数
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

// Read 方法
func (a *alphaReader) Read(p []byte) (int, error) {
	// 当前位置 >= 字符串长度 说明已经读取到结尾 返回 EOF
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	// x 是剩余未读取的长度
	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		// 剩余长度超过缓冲区大小，说明本次可完全填满缓冲区
		bound = len(p)
	} else if x < len(p) {
		// 剩余长度小于缓冲区大小，使用剩余长度输出，缓冲区不补满
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

func main() {
	// tem := T{A: 1, B: 2}
	// fmt.Println(tem.A)
	fmt.Println(time.Now().Format("15:04:05\n"))
	read := newAlphaReader("Hello! It's 9am, where is the sun?")
	//p := make([]byte, 4)
	// p := make([]byte, 4)
	// for {
	// 	n, err := read.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Print(string(p[:n]))
	// }

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(read)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)

	return

	//

}
