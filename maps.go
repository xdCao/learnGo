package main

import (
	"fmt"
)

func main() {

	m := map[string]string{
		"name":    "xdCao",
		"course":  "golang",
		"site":    "mooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) /*m2 = empty map*/

	var m3 map[string]int /*m3 = nil*/

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	/*判断元素是否存在*/
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	courseName1, ok1 := m["couse"]
	fmt.Println(courseName1, ok1)

	/*删除元素*/
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(nonRepeating("abcabcbb"))

}

func nonRepeating(s string) int {

	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i

	}

	return maxLength

}
