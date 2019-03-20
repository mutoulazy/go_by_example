package main

import (
	"fmt"
	"sort"
)

/*
	自定义排序方法 实现Len Swap Less 三个接口
*/
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func testCustomSort() {
	fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}

func main() {
	// 字符排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 数字排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 判断是否已经完成排序
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)

	testCustomSort()
}