package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

/**
数据组的定义
*/

func ArrayStatement() {
	var a [3]int
	fmt.Println(a)

	var b = [...]int{1, 2, 3}
	fmt.Println(b)

	var c = [...]int{2: 3} // 从索引为2开始设置值为3
	fmt.Println(c)

	var d = [...]int{1, 2, 2: 3, 6}
	fmt.Println(d)
	for i, v := range d {
		fmt.Printf("index %d, value is %d\n", i, v)
	}
	for index := 0; index < len(d); index++ {
		fmt.Printf("Value is: %d\n", d[index])
	}
}

/**
数组传值:
Go 语言中数组是值语义。一个数组变量即表示整个数组，它并不是隐式的指向第一个元素的指针（比如 C 语言的数组），而是一个完整的值。当一个数组变量被赋值或者被传递的时候，
实际上会复制整个数组。如果数组较大的话，数组的赋值也会有较大的开销。为了避免复制数组带来的开销，可以传递一个指向数组的指针，但是数组指针并不是数组。



=是赋值语句，但是此处的变量必须提前声明好类型，或者在声明的同时赋值。
:=是声明类型并赋值，自动匹配变量类型，省去了声明变量类型操作，也就是你可以不声明，在赋值使用的时候多加个冒号就可以了。
*/

func ArrayTransmission() {
	var a = [...]int{1, 2, 2: 3, 6}
	fmt.Println(a)
	b := a
	fmt.Println(b)
	var c = &a
	fmt.Println(c[0])
}

/**
数组不仅仅可以用于数值类型，还可以定义字符串数组、结构体数组、函数数组、接口数组、管道数组等等：
*/

func TypeOfArray() {
	// 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	fmt.Println(s1, s2, s3)

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1, line2, line3)

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Println(decoder1, decoder2)

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	fmt.Println(unknown1, unknown2)
	// 管道数组
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("test_channel")
		c1 <- [0]int{}
	}()
	<-c1

}

func main() {
	//ArrayStatement()
	//ArrayTransmission()
	TypeOfArray()
}
