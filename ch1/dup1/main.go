// dup的第一个版本打印标准输入中多次出现的行，以重复次数开头。
// 该程序将引入if语句，map数据类型以及bufio包。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
