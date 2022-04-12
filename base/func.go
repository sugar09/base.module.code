package base

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

// RunRootPath 获取程序运行目录
var RunRootPath = func() (string, error) {
	f, e := os.Executable()
	//获取执行文件失败
	if e != nil {
		return "", e
	}
	p := filepath.Dir(f)
	return p, nil
}

// FileMd5 文件md5值
var FileMd5 = func(path string) string {
	f, e := os.Open(path)
	if e != nil {
		return ""
	}
	defer f.Close()
	m := md5.New()
	io.Copy(m, f)
	return hex.EncodeToString(m.Sum(nil))
}

// PathExists 检测文件路径是否存在
var PathExists = func(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// MakeDirs 创建目录；多层创建
var MakeDirs = func(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

// 检测文件夹是否存在；不存在创建
var CheckDirExistsWithCreate = func(p string) bool {
	if !PathExists(p) {
		return MakeDirs(p)
	}
	return true
}
