package main

func reverseWords(s string) string {
	var res = ""
	var start = -1
	for idx, v := range s { // 最后一个单词会有问题
		if string(v) == " " {
			for i := idx - 1; i > start; i-- {
				res += string(s[i])
			}
			start = idx
		}
	}
	return res
}
