package base

import (
	"fmt"
	"testing"
)

func TestInterfaceToString(t *testing.T) {
	var data interface{}

	data = nil
	fmt.Println(InterfaceToString(data))
}
