// Step7
// ポインタを利用してみましょう．
// あなたの年齢を格納したint型の変数 age を main 関数の中で定義してください．
// その後，引数にageのアドレスを受け取ってageを+1する，birthDay関数を定義してください．
// ※birthDay関数は値をreturnしてはいけません．
package main

import (
	"fmt"
)

func birthDay(age *int) {
	*age += 1
	fmt.Println(*age)
}

func main() {
	var age int = 20
	fmt.Println(age)
	birthDay(&age)
	//fmt.Println(birthDay(age_p))
}
