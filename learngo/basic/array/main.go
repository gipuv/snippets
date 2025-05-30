package main

import (
	"fmt"
	"reflect"
)

func base() {
	// 数组声明
	var a [5]int

	// 初始化
	// 1.逐个赋值
	a[0] = 1
	a[1] = 2

	// 求和
	sum := 0
	for _, val := range a {
		sum += val
	}

	fmt.Println("a数组和：", sum)

	// 声明时初始化
	b := [5]int{1, 2, 3, 4, 5}

	// 数组元素访问
	fmt.Println(b[1])

	// 最大值、最小值
	max := b[0]
	min := b[0]

	for _, v := range b {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	fmt.Println("最大值:", max, "最小值:", min)

	// 反转数组
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1] = b[len(b)-1], b[i]
	}

	fmt.Println("b反转：", b)

	// 自动推断长度
	c := [...]int{10, 20, 30}

	// for i := 0; i < len(c); i++ {
	for i := range len(c) { // 同上
		fmt.Println(c[i])
	}

	for index, value := range c {
		fmt.Printf("Index %d = %d\n", index, value)
	}

	// 数组是值类型，赋值或传参时会复制整个数组

	d := c
	d[0] = 100

	fmt.Println("c:", c) // c: [10, 20, 30]
	fmt.Println("d:", d) // d: [100, 20, 30]

	// 如果你想引用传递，使用指针或切片（slice）
	e := &c
	e[0] = 1000

	fmt.Println("c:", c)                    // c: [1000 20 30]
	fmt.Println("d:", d)                    // d: [100 20 30]
	fmt.Println("e:", e, reflect.TypeOf(e)) // e: &[1000 20 30] *[3]int
}

// 修改原数组
func bubbleSort(arr *[6]int) {
	n := len(arr)

	// 冒泡排序
	for i := range n - 1 {
		for j := range n - 1 - i {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 不修改原数组
func bubbleSortCopy(arr [6]int) [6]int {
	// 复制原数组
	sorted := arr

	n := len(sorted)

	// 冒泡排序
	for i := range n - 1 {
		for j := range n - 1 - i {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}

func main() {
	base()

	// 练习
	// 创建一个 [10]int 数组，初始化为 1~10，输出所有偶数。
	a := [10]int{}
	for i := range 10 {
		a[i] = i + 1 // 从 1 到 10
	}

	evenCount := 0
	for i := range a {
		if a[i]%2 == 0 {
			evenCount++
		}
	}

	fmt.Println("偶数个数：", evenCount)

	// 统计 [6]int{1, 1, 2, 2, 3, 3} 中某个值出现的次数。
	b := [6]int{1, 1, 2, 2, 3, 3}
	target := 1
	count := 0

	for _, v := range b {
		if v == target {
			count++
		}
	}

	fmt.Printf("值 %d 出现了 %d 次\n", target, count)

	// 输入一个 [5]int 数组，求平均值。
	c := [5]int{1, 2, 3, 4, 5}
	sum := 0
	for _, v := range c {
		sum += v
	}
	avg := float64(sum) / float64(len(c))
	fmt.Printf("c 平均值：%.2f\n", avg)

	data := [6]int{5, 2, 9, 1, 3, 7}
	fmt.Println("排序前：", data)
	bubbleSort(&data)
	fmt.Println("排序后：", data)

	original := [6]int{5, 2, 9, 1, 3, 7}
	fmt.Println("原数组：", original)

	sorted := bubbleSortCopy(original)
	fmt.Println("排序后：", sorted)
	fmt.Println("原数组依旧：", original) // 验证未修改
}
