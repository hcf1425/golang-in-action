package main

import "fmt"

type TZ int

func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}

func (tz *TZ) Increase(num int) {
	*tz += TZ(num) // 一定要进行将int类型转换为TZ类型才可进行计算。
}
