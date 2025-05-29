package main

import (
	"fmt"
	"reflect"
)

// & —— 取地址符
// * —— 指针解引用（取值）
// * --	声明指针类型 var p *int

type Vertex struct {
	X int
	Y int
}

func basePointer() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值 42
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值 21

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值 73

	var a int = 100
	var varP *int = &a // p 是指向 a 的指针

	// 类型
	fmt.Println(reflect.TypeOf(varP))
	fmt.Printf("%T\n", a) // int

	fmt.Println(*varP, varP) // 100
}

func structPointer() {
	v := Vertex{1, 2}
	vp := &v
	fmt.Println("vp 类型：", reflect.TypeOf(vp)) // *main.Vertex
	vp.X = 1e9
	fmt.Println(v)
}

func structLiteral() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2} // 创建一个 *main.Vertex 类型的结构体（指针）
	)

	fmt.Println(v1, v2, v3, p)
}

func main() {
	// 指针
	basePointer()
	// 42
	// 21
	// 73
	// *int
	// int
	// 100 0xc00000a138

	// 结构体指针
	structPointer()
	// vp 类型： *main.Vertex
	// {1000000000 2}

	// 结构体字面量
	structLiteral()
	// {1 2} {1 0} {0 0} &{1 2}
}
