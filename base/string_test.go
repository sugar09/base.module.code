package base

import (
	"fmt"
	"testing"
)

func TestStringContains(t *testing.T) {
	str := "有些测试需要一定的启动和初始化时间"
	sub := "初始化"

	fmt.Println(StringContains(str, sub))
}
