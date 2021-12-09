// Step3
// 引数に名前(文字列)を受け取って，"Hello,名前"のような文字列を出力して
// returnする，sayHelloという関数を制作してください．
package main

import (
	"fmt"
)

func sayHello(name string) string {
	a := "Hello," + name
	fmt.Println(a)
	return a
}
func main() {
	name := "Go"
	name2 := "Rust"
	name3 := "Python"
	sayHello(name)
	sayHello(name2)
	sayHello(name3)
}
