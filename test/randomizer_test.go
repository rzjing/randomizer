/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         randomizer_test.go
@ Create Time:  2020/4/15 15:48
@ Software:     GoLand
*/

package main

import (
	"math"
	"randomizer/handler"
	"testing"
)

func TestRInt64(t *testing.T) {
	tests := []struct{ min, max int64 }{
		{0, 0},
		{0, 9},
		{9, 0},
		{math.MinInt64, math.MaxInt64},
	}

	for _, test := range tests {
		if actual := handler.RInt64(test.min, test.max); actual <= test.min || actual >= test.max {
			t.Errorf("RInt64(%d, %d); "+"expected %d", test.min, test.max, actual)
		}
	}
}
