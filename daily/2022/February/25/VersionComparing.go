package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.001", "1.0001"))
}

func compareVersion(version1 string, version2 string) int {

	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	l1, l2 := len(v1), len(v2)
	var n int
	if l1 > l2 {
		n = l1
	} else {
		n = l2
	}

	for i := 0; i < n; i++ {
		lv, rv := 0, 0

		if i < l1 {
			lv, _ = strconv.Atoi(v1[i])
		}

		if i < l2 {
			rv, _ = strconv.Atoi(v2[i])
		}

		if lv > rv {
			return 1
		}

		if rv > lv {
			return -1
		}
	}

	return 0

}
