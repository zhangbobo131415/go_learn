package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// tem := T{A: 1, B: 2}
	// fmt.Println(tem.A)
	fmt.Println(time.Now().Format("15:04:05\n"))
	go spinner(200 * time.Millisecond)
	const n = 10
	_ = fib(n) // slow
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	time.Sleep(3 * time.Second)
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
	fmt.Println("jejj")

	return

	//

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
