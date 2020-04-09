/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         randomizer.go
@ Create Time:  2020/4/8 19:00
@ Software:     GoLand
*/

package handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var randomSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

func RInt64(min, max int64) int64 {
	if min >= max || min < 0 || max == 0 {
		return max
	}
	return randomSeed.Int63n(max-min) + min
}

func RString(length int) string {
	buffer := make([]byte, length)
	for i := range buffer {
		buffer[i] = charset[randomSeed.Intn(len(charset))]
	}
	return string(buffer)
}

func RMac() string {
	buffer := make([]byte, 6)
	for i := range buffer {
		buffer[i] = charset[randomSeed.Intn(len(charset))]
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buffer[0], buffer[1], buffer[2], buffer[3], buffer[4], buffer[5])
}

func RIPv4() string {
	buffer := make([]string, 4)
	for i := range buffer {
		buffer[i] = strconv.Itoa(randomSeed.Intn(255))
	}
	return fmt.Sprintf("%s.%s.%s.%s", buffer[0], buffer[1], buffer[2], buffer[3])
}

func RProvince() string {
	return province[randomSeed.Intn(len(province))]
}

func RCity() string {
	return city[randomSeed.Intn(len(city))]
}
