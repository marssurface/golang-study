// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
//

package main

import (
	"fmt"
	"os"
)

func main()  {
	var s, sep string

	// 版本 1
	//for i := 0; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}

	//版本 2
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)


	// 版本 3
	//fmt.Println(strings.Join(os.Args[0:], " "))

}