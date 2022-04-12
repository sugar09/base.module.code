package logs

import (
	"encoding/json"
	"fmt"
	"github.com/sugar09/base.module.code/base"
	"testing"
)

func TestBeegoLogsConf(t *testing.T) {
	conf := &base.BeegoLogsConf{
		LogsDaily:    false,
		LogsFileName: "log",
	}

	b, _ := json.Marshal(conf)
	fmt.Println(string(b))
}
