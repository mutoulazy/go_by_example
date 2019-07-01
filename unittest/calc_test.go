package main

import (
	"fmt"
	"testing"
)

func testAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add r error %d", r)
	}
	t.Logf("add success")
}

func testSub(t *testing.T) {
	r := sub(4, 2)
	if r != 2 {
		t.Fatalf("sub r error %d", r)
	}
	t.Logf("sub success")
}

func TestAll(t *testing.T) {
	t.Run("TestAdd", testAdd)
	t.Run("TestSub", testSub)
}

func TestMain(m *testing.M) {
	fmt.Println("test begin... 可以做一些初始化操作")
	m.Run()
}

func testA(n int) int {
	return n
}

// go的性能测试 go test -bench='.'
func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testA(n)
	}
}
