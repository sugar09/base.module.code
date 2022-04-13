package base

// StringContains 检测str是否包含sub
func StringContains(str, sub string) int {
	lr := len(str)
	lb := len(sub)

	//如果sub长度大于str，或sub长度等于0；
	if lr < lb || lb == 0 {
		return -1
	}

	//如果str长度等于sub长度
	if lr == lb {
		if str == sub {
			return 0
		}
		return -1
	}

	next := getNext(sub)
	j := 0
	for i := 0; i < lr; i++ {
		for j > 0 && str[i] != sub[j] {
			j = next[j-1] + 1
		}
		if str[i] == sub[j] {
			j++
		}
		if j == lb {
			return i - j + 1
		}
	}
	return -1
}

func getNext(s string) []int {
	var next = make([]int, len(s))
	next[0] = -1
	k := -1
	for i := 1; i < len(s); i++ {
		for k != -1 && s[k+1] != s[i] {
			k = next[k]
		}
		if s[k+1] == s[i] {
			k++
		}
		next[i] = k
	}
	return next
}
