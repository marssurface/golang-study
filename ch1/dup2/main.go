// 很多程序要么从标准输入中读取数据，如上面的例子所示，要么从一系列具名文件中读取数据。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// 没有输入 命令参数 则从标准输入中读取数据
		countLines(os.Stdin, counts)
	} else {
		// 从文件读入
		for _, arg := range files {
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File）
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup 2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
