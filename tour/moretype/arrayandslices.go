package moretype

import "fmt"

// var a [10]int
// 将变量 a 声明为拥有10个整数的数组
// 数组的长度是其类型的一部分，因此数组不能改变大小

func arrayDef() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// []T 表示一个元素类型为 T 的切片
// 切边可以通过两个下标来界定：a[low : high]
// 半开区间，包括下标为 low 的元素，不包括下标为 high 的元素
// 切片下界的默认值为 0，上界则是该切片的长度：a[0:10] <=> a[:10] <=> a[0:] <=> a[:]

// 切片就像数组的引用
// 切片并不存储任何数据，它只是描述了底层数组中的一段
// 更改切片的元素会修改其底层数组中对应的元素

// [3]bool{true, true, false} // 数组
// []bool{true, true, false} // 切片

func sliceDef() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println()
	sliceModifyElement()
	fmt.Println()

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	f := []struct {
		x int
		y bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(f)
	fmt.Println()
	sliceDefault()
	fmt.Println()
	sliceLenAndCap()
	fmt.Println()
	sliceNil()
	fmt.Println()
	sliceMake()
}

func sliceModifyElement() {
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
	fmt.Println(a, b, names)
}

func sliceDefault() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	fmt.Printf("len: %v, caps: %v\n", len(s), cap(s))

	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("len: %v, caps: %v\n", len(s), cap(s))

	s = s[:2]
	fmt.Println(s)
	fmt.Printf("len: %v, caps: %v\n", len(s), cap(s))

	s = s[1:]
	fmt.Println(s)
	fmt.Printf("len: %v, caps: %v\n", len(s), cap(s))
}

func sliceLenAndCap() {
	// 切片的长度：切片所包含元素的个数
	// 切片的容量：从切片的第一个元素开始数，到其底层数组元素末尾的个数
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, s = %v\n", len(s), cap(s), s)
}

func sliceNil() {
	// 切片的零值是 nil
	// nil 切片的长度和容量为0，且没有底层数组
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
	var b [10]string
	fmt.Printf("%q\n", b)
}

func sliceMake() {
	// 长度、容量都为5
	a := make([]int, 5)
	printSlice2("a", a)

	// 长度为0，容量为5
	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:5]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
