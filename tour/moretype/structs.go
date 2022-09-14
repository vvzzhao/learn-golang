package moretype

import "fmt"

type Vertex struct {
	X int
	Y int
}

/*
type Vertex struct {
	X,Y int
}
*/

var (
	v1 = Vertex{1, 2} // 创建一个 Vertex 类型的结构体
	// 使用 Name: 语法可以仅列出部分字段（字段名的顺序无关）
	v2 = Vertex{X: 1}        // Y:0 被隐式地赋予
	v3 = Vertex{}            // X:0 Y:0
	p  = &Vertex{Y: 1, X: 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func structDef() {
	v := Vertex{1, 2}
	cv := v
	v.Y = 3
	fmt.Printf("v = %v, cv = %v\n", v, cv)
	fmt.Println()
	fmt.Println(v1, v2, v3, p)
}
