package search

import (
	"strconv"
	"strings"
)

/**
BM22 比较版本号

*/

func compare(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	p := 0
	for p < len(v1) || p < len(v2) {
		t1, t2 := 0, 0
		if p < len(v1) {
			t1, _ = strconv.Atoi(v1[p])
		}
		if p < len(v2) {
			t2, _ = strconv.Atoi(v2[p])
		}
		if t1 > t2 {
			return 1
		} else if t1 < t2 {
			return -1
		}
		p++
	}
	return 0
}
