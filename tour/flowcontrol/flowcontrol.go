package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	_for()
	_if()
	_switch()
	_defer()
}

func _for() {
	fmt.Println("=========  for  =========")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("0 + 1 + 2 + ... + 9 =", sum)
	fmt.Println("-------------------------")
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println("sum =", sum)
	fmt.Println("-------------------------")
	sum = 1
	// go 中的 "while"
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum =", sum)
	fmt.Println("-------------------------")
	sum = 1
	// 无限循环
	for {
		if sum >= 1000 {
			break
		}
		sum += sum
	}
	fmt.Println("sum =", sum)
	fmt.Println("=========================")
}

func _if() {
	fmt.Println("=========   if   =========")
	fmt.Println("sqrt(2) =", sqrt(2))
	fmt.Println("sqrt(0) =", sqrt(0))
	fmt.Println("sqrt(-4) =", sqrt(-4))
	fmt.Println("pow(3, 2, 10) =", pow(3, 2, 10))
	fmt.Println("pow(3, 3, 10) =", pow(3, 3, 20))
	fmt.Println("==========================")
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else if 0 == x {
		fmt.Println("else if x == 0")
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("else printf: %g >= %g\n", v, lim)
	}
	return lim
}

func _switch() {
	fmt.Println("========= switch =========")
	printRunOn()
	fmt.Println()
	printIntervalSaturday()
	fmt.Println()
	seeGood()
	fmt.Println("==========================")
}

func printRunOn() {
	var runOn string
	switch os := runtime.GOOS; os {
	case "darwin":
		runOn = "OS X"
	case "linux":
		runOn = "Linux"
	default:
		runOn = os
	}
	fmt.Printf("Go runs on %s.\n", runOn)
}

func printIntervalSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func seeGood() {
	hour := time.Now().Hour()
	// 没有条件的 switch 同 switch true 一样。
	// 这种形式能将一长串 if-then-else 写得更加清晰。
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// https://blog.go-zh.org/defer-panic-and-recover
func _defer() {
	// defer 语句会将函数推迟到外层函数返回之后执行。
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	fmt.Println("=========  defer  =========")
	deferPrintHelloWorld()
	fmt.Println()
	printNineToZero()
	fmt.Println("===========================")
}

func printNineToZero() {
	fmt.Print("nine to zero: ")
	for i := 0; i < 10; i++ {
		if 0 == i {
			defer fmt.Println(i)
		} else {
			defer fmt.Printf("%d ", i)
		}
	}
}

func deferPrintHelloWorld() {
	defer fmt.Println(" world!")
	fmt.Print("hello")
}
