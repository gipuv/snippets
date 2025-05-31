package main

import (
	"fmt"
	"reflect"
	"strings"
)

func basic() {
	arr := []int{1, 2, 3, 4, 5, 6}

	var s = arr[2:]

	fmt.Println(s)
}

func edit() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}

func lenAndCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 长度为0
	s = s[:0] // 不删除底层数组，只是修改切片的视图。
	printSlice(s)

	// 扩展长度
	s = s[:4]
	printSlice(s)

	// 舍弃两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	var nilS []int
	printSlice(nilS)
	fmt.Println(reflect.TypeOf(nilS))

	// 用make创建切片
	a := make([]int, 5) // 长度是5 容量是5
	printSlice(a)

	b := make([]int, 0, 5) // 长度0 容量5
	printSlice(b)

	// 切片 [:x] 只改 len，不会改 cap
	c := b[:2]
	printSlice(c)
	// cap 是基于切片创建时的底层数组决定的，除非重新分配数组
	b = append(b, 1, 2, 3, 4, 5) // len=5 cap=5，刚好填满
	b = append(b, 6)             // 超过 cap=5，触发重新分配
	printSlice(b)

	c = b[:2]
	printSlice(c)

	c1 := c[0:len(b)] // len(b) 底层合理 / len(c) 合理&安全
	printSlice(c1)

	// 约束 1
	if len(b) <= cap(c) {
		d := c[0:len(b)]
		fmt.Println("约束1 d:")
		printSlice(d)
	} else {
		fmt.Println("不能从 c 中取出 len(b) 的长度，超出 cap")
	}

	// 约束 2
	d := c[0:cap(c)] // 最多取到 c 的容量上限
	fmt.Println("约束2 d:")
	printSlice(d)
}

func sliceSlice() {
	// 创建一个 3x3 的二维切片，表示井字棋的棋盘
	// 每一行是一个切片，整个棋盘是由三行组成的切片的切片（二维切片）
	board := [][]string{
		[]string{"_", "_", "_"}, // []string 可以省略
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	for i := range len(board) {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	s = append(s, 5)
	printSlice(s)

	// 当前容量6，长度6，已满，需要扩容
	// 扩容后容量可能变为12（通常是之前容量的2倍）
	s = append(s, 6, 7)
	printSlice(s)
}

func rangeSlice() {
	// pow 是一个整数切片，包含2的幂次方
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for _, v := range pow {
		v *= 2
		fmt.Printf("修改为 %d（副本）\n", v)
	}

	printSlice(pow)

	for i := range pow {
		pow[i] *= 2 // 修改原切片
	}

	printSlice(pow)
}

func main() {
	basic()
	edit()
	lenAndCap()
	makeSlice()
	sliceSlice()
	appendSlice()
	rangeSlice()
}
