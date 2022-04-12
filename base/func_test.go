package base

import (
	"fmt"
	"testing"
)

// TestRunRootPath 测试获取执行目录
func TestRunRootPath(t *testing.T) {
	path, err := RunRootPath()
	fmt.Println(path, err)
}

func TestFileMd5(t *testing.T) {
	fmt.Println(FileMd5("./api.go"))
}

func TestPathExists(t *testing.T) {
	fmt.Println(PathExists("./api.gox"))
}