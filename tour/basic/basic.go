// 程序从 main 包开始运行（package main）
package main

// 按照约定，包名与导入路径的最后一个元素一致。如，math/rand 包中的源码均已 package rand 语句开始

/*
import "fmt"
import "math/rand"
// or
import (
	"fmt"
	"math/rand"
)
*/
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// 如果一个名字以大写字母开头，那么则表示它已导出。
	// 导出方法如：Println，Intn
	// 导出变量如：Pi
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("math.Pi =", math.Pi)

	fmt.Println()

	// add 函数
	fmt.Println("add(1, 2) =", add(1, 2))
	x, y := 3, 4
	// 多值返回
	var sx, sy int = swap(x, y)
	fmt.Printf("swap(%v, %v) = %v, %v\n", x, y, sx, sy)
	// 命名返回值
	sx, sy = swap2(x, y)
	fmt.Printf("swap2(%v, %v) = %v, %v\n", x, y, sx, sy)

	fmt.Println()
	varDef()
	fmt.Println()
	basicTypes()
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y int) (a, b int) {
	a, b = y, x
	return
}

func varDef() {
	fmt.Println("=========  变量  =========")
	// 变量
	var a, b, c bool
	fmt.Println("a, b, c =", a, b, c)

	var d, e int = 1, 2
	var f, g = true, "no"
	fmt.Println("d, e, f, g =", d, e, f, g)

	h, i := 123, false
	fmt.Println("h, i =", h, i)

	fmt.Println("==========================")
}

func basicTypes() {
	fmt.Println("========= 基本类型 =========")
	a, b, i32 := true, "str", 123
	i8 := int8(123)
	var i16 int16 = 123
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	f32 := float32(123.123)
	f64 := 123.123
	var c64 complex64 = 123 + 3i //c64 := complex64(123 + 3i)
	c128 := 123 + 3i

	printTypeAndValue("bool",a)
	printTypeAndValue("string",b)
	printTypeAndValue("int",i32)
	printTypeAndValue("int8",i8)
	printTypeAndValue("int16",i16)
	printTypeAndValue("float32",f32)
	printTypeAndValue("float64",f64)
	printTypeAndValue("complex64",c64)
	printTypeAndValue("complex128",c128)
	fmt.Println("==========================")
}

func printTypeAndValue(expectType string, a interface{}) {
	fmt.Printf("Expect Type: %s ,Type: %T Value: %v\n", expectType, a, a)
}