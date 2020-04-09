/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         main.go
@ Create Time:  2020/4/8 18:59
@ Software:     GoLand
*/

package main

import (
	"fmt"
	"randomizer/handler"
)

func main() {
	// 随机 数字
	fmt.Println(handler.RInt64(0, 100))
	// 随机 字符串
	fmt.Println(handler.RString(8))
	// 随机 MAC 地址
	fmt.Println(handler.RMac())
	// 随机 IPv4 地址
	fmt.Println(handler.RIPv4())
	// 随机 省份
	fmt.Println(handler.RProvince())
	// 随机 城市
	fmt.Println(handler.RCity())
}
