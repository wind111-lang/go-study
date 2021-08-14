// Step6
// Student構造体を定義して，そのメンバーを表示するプログラムを制作してください．
// 構造体のメンバ変数は，以下のように定義してください．(値は適当に決めてください)
// 	name: "Taro",
// 	age:  20,
// 	number: "K99999"
package main

import "fmt"

type list struct {
	name   string
	age    int
	number string
}

func main() {
	var profile list
	profile.name = "wind111"
	profile.age = 20
	profile.number = "K20076"

	fmt.Println(profile)
}
