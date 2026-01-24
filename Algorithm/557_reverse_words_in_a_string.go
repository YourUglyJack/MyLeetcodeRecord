package main

import "fmt"

func reverseWords(s string) string {
	var left, right = 0, 0
	var b = []byte(s)
	var l = len(b)
	//fmt.Println(b)
	for i := 0; i < l; i++ {
		if b[i] != ' ' {
			right++
		} else {
			// fmt.Println(left, right, l, b[i])
			var des = (right - left) / 2
			for j := 0; j < des; j++ {
				var tmp = b[left+j]
				b[left+j] = b[right-j-1]
				b[right-1-j] = tmp
			}
			left = i + 1
			right++
			//fmt.Println(b, string(b))
		}
		if right == l {
			// fmt.Println(left, right, b[i])
			var des = (right - left) / 2
			for j := 0; j < des; j++ {
				var tmp = b[left+j]
				b[left+j] = b[right-j-1]
				b[right-1-j] = tmp
			}
		}
	}
	return string(b)
}

func main() {
	var s = "Let's take LeetCode contest"
	var res = reverseWords(s)
	fmt.Println(res)
}
