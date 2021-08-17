// Step8
// Step6で制作したStudent構造体に，以下のようなメソッドを追加してください．
// 	- getName() : 名前を返すメソッド
// 	- getAge() : 年齢を返すメソッド
// 	- getNumber() : 学籍番号を返すメソッド
// 	- selfIntroduction() : 自己紹介文を返すメソッド
//  - birthDay() : 年齢を+1するメソッド
package main

import (
	"fmt"
	"strconv"
)

type student struct {
	name   string
	age    int
	number string
}

func (name student) getName() string {
	return name.name
}

func (age student) getAge() int {
	return age.age
}

func (num student) getNumber() string {
	return num.number
}

func (student student) selfIntroduction() string {
	introduction := "Hi! I'm " + student.getName() + ". " + strconv.Itoa(student.getAge()) + " years old. My number is " + student.getNumber() + "."
	return introduction
}

func (age *student) birthDay() int {
	age.age = age.getAge() + 1
	return age.age
}

func main() {
	var student student

	student.name = "wind111"
	student.age = 20
	student.number = "K20076"

	// fmt.Println(getName(student))
	// fmt.Println(getAge(student))
	// fmt.Println(getNumber(student))
	fmt.Println(student.selfIntroduction())
	fmt.Println(student.birthDay())
}
