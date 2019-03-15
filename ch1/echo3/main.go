// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
//（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	var s, sep string

	//t := time.Now().UnixNano()
	////fmt.Println(t.Unix()) // 时间戳 秒级
	//fmt.Println(t.UnixNano()) // us
	//time.Sleep(time.Duration(2) * time.Second)
	//s := time.Now()
	//fmt.Println(s.UnixNano())

	var arr [10000]string

	str := "fffdfsafsadfasfasfd"
	for j := 0; j < 10000; j++ {
		arr[j] = str
	}

	a := time.Now().UnixNano()
	// 版本 1
	for i := 0; i < len(arr); i++ {
		s += sep + arr[i]
		sep = " "
	}

	b := time.Now().UnixNano()

	//版本 2
	for _, arg := range arr[0:] {
		s += sep + arg
		sep = " "
	}
    // for range 效率 比 for 低 因为 for range 是值复制  for 是 索引
	c := time.Now().UnixNano()

	//版本 3
	s =  strings.Join(os.Args[0:], " ")

	d := time.Now().UnixNano()

	fmt.Println(b - a)
	fmt.Println(c - b)
	fmt.Println(d - c)


}
