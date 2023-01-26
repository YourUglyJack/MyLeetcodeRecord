package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	t := []bool{false, true}
	return t[version-1]
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := (right-left)/2 + left
		if isBadVersion(mid) == true {
			if isBadVersion(mid-1) == false {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return left // 2ms
}

func firstBadVersionV2(n int) int {
	left := 1
	right := n - 1
	for left <= right {
		mid := (right-left)/2 + left
		if isBadVersion(mid) == true {
			if isBadVersion(mid-1) == false {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return left // 1 ms
}

func main() {
	res := firstBadVersion(2)
	fmt.Println(res)
}
