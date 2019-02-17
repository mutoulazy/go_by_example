package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus3(a, b, c int) int {
	return a + b + c
}

// 多返回值
func vals(a, b int) (int, int) {
	return b, a
}

// 变参函数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包(匿名函数)
func intseq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

// 递归函数
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plus3(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals(1, 2)
	fmt.Println(a, b)

	_, c := vals(1, 2)
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	nextInt := intseq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intseq()
	fmt.Println(newInts())

	fmt.Println(fact(7))
}
