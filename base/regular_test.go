package base

import "testing"

// TestCheckInt 测试验证int
func TestCheckInt(t *testing.T) {
	//测试案例
	cases := []struct {
		value   int
		wantErr bool
	}{
		{
			value:   9,
			wantErr: false,
		},
		{
			value:   0,
			wantErr: false,
		},
		{
			value:   07,
			wantErr: false,
		},
		{
			value:   10,
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckInt(val.value)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%d, bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%d, bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckInts 测试验证多个int逗号分割
func TestCheckInts(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "9",
			wantErr: false,
		},
		{
			value:   "1,23",
			wantErr: false,
		},
		{
			value:   ",11",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckInts(val.value)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s, bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s, bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckWStr 验证数字、字母组成的字符串
func TestCheckWStr(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "9",
			wantErr: false,
		},
		{
			value:   "1,23",
			wantErr: false,
		},
		{
			value:   "Wxas16",
			wantErr: false,
		},
		{
			value:   "231qwe",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckWStr(val.value)
		t.Log(val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s, bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s, bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckWStrWithLen 验证数字、字母组成的指定长度的字符串
func TestCheckWStrWithLen(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		len     int
		wantErr bool
	}{
		{
			value:   "912",
			len:     3,
			wantErr: false,
		},
		{
			value:   "423",
			len:     6,
			wantErr: false,
		},
		{
			value:   "Wxas16",
			len:     6,
			wantErr: false,
		},
		{
			value:   "231qwe",
			len:     6,
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckWStrWithLen(val.value, val.len)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s len %d bool:%t", val.value, val.len, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s len %d bool:%t", val.value, val.len, b)
			}
		}
	}
}

// TestCheckWStrWithRan 验证数字、字母组成的指定范围长度的字符串
func TestCheckWStrWithRan(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		min     int
		max     int
		wantErr bool
	}{
		{
			value:   "912",
			min:     2,
			max:     4,
			wantErr: false,
		},
		{
			value:   "423",
			min:     1,
			max:     2,
			wantErr: false,
		},
		{
			value:   "Wxas16",
			min:     4,
			max:     6,
			wantErr: false,
		},
		{
			value:   "231qwe",
			min:     6,
			max:     6,
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckWStrWithRan(val.value, val.min, val.max)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s min %d max %d bool:%t", val.value, val.min, val.max, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s min %d max %d bool:%t", val.value, val.min, val.max, b)
			}
		}
	}
}

// TestCheckUniqueStr 验证数字、字母、-组成的字符串
func TestCheckUniqueStr(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "9",
			wantErr: false,
		},
		{
			value:   "1,23",
			wantErr: false,
		},
		{
			value:   "Wxa-+",
			wantErr: false,
		},
		{
			value:   "231qwe",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckUniqueStr(val.value)
		t.Log(val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s, bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s, bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckUniqueStrWithLen 验证数字、字母、-组成的指定长度的字符串
func TestCheckUniqueStrWithLen(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		len     int
		wantErr bool
	}{
		{
			value:   "912",
			len:     3,
			wantErr: false,
		},
		{
			value:   "423",
			len:     6,
			wantErr: false,
		},
		{
			value:   "Wxas16",
			len:     6,
			wantErr: false,
		},
		{
			value:   "231qwe",
			len:     6,
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckUniqueStrWithLen(val.value, val.len)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s len %d bool:%t", val.value, val.len, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s len %d bool:%t", val.value, val.len, b)
			}
		}
	}
}

// TestCheckUniqueStrWithRan 验证数字、字母、-组成的指定范围长度的字符串
func TestCheckUniqueStrWithRan(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		min     int
		max     int
		wantErr bool
	}{
		{
			value:   "912",
			min:     2,
			max:     4,
			wantErr: false,
		},
		{
			value:   "423",
			min:     1,
			max:     2,
			wantErr: false,
		},
		{
			value:   "Wxas16",
			min:     4,
			max:     6,
			wantErr: false,
		},
		{
			value:   "231qwe",
			min:     6,
			max:     6,
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckUniqueStrWithRan(val.value, val.min, val.max)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s min %d max %d bool:%t", val.value, val.min, val.max, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s min %d max %d bool:%t", val.value, val.min, val.max, b)
			}
		}
	}
}

// TestCheckCnEnNumStr 验证常见的中英和其他文字符组成的字符串
func TestCheckCnEnNumStr(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "刘江13",
			wantErr: false,
		},
		{
			value:   "刘江、@#￥%、/\\|",
			wantErr: false,
		},
		{
			value:   "1236445",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckCnEnNumStr(val.value)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckEmail 验证邮箱
func TestCheckEmail(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "dwqeda@aaa.con",
			wantErr: false,
		},
		{
			value:   "2446422230@qq.com",
			wantErr: false,
		},
		{
			value:   "216@168.com",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckEmail(val.value)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckEmail 验证手机号
func TestCheckPhone(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "18379136664",
			wantErr: false,
		},
		{
			value:   "1532649555",
			wantErr: false,
		},
		{
			value:   "19202222001",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckPhone(val.value)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckTimestamp 验证时间戳
func TestCheckTimestamp(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "18379136664",
			wantErr: false,
		},
		{
			value:   "1532649555",
			wantErr: false,
		},
		{
			value:   "19202222001",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckTimestamp(val.value)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s bool:%t", val.value, b)
			}
		}
	}
}

// TestCheckDateTime 验证日期时间
func TestCheckDateTime(t *testing.T) {
	//测试案例
	cases := []struct {
		value   string
		wantErr bool
	}{
		{
			value:   "2021-11-10 15:30:11",
			wantErr: false,
		},
		{
			value:   "2021-11-10 15:30:112",
			wantErr: false,
		},
		{
			value:   "2021-11-10 15:3",
			wantErr: false,
		},
	}

	for _, val := range cases {
		b := CheckDateTime(val.value)
		t.Logf("%#v\n", val)
		if val.wantErr {
			if b {
				t.Errorf("想出错，却没出错！value：%s bool:%t", val.value, b)
			}
		} else {
			if !b {
				t.Errorf("不想出错，却出错了！value：%s bool:%t", val.value, b)
			}
		}
	}
}
