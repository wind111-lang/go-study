// Step4
// 引数が3の倍数ならFizz,5の倍数ならBuzz,
// 3の倍数かつ5の倍数ならFizzBuzz,
// どれでもない場合は引数をそのまま出力するFizzBuzz関数を制作し，
// 引数に1~100を与えて実行してください．
package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(num int) string { //intで代入して、結果はstringで返す
	if num%3 == 0 && num%5 == 0 {
		return "FizzBuzz"
	} else if num%3 == 0 {
		return "Fizz"
	} else if num%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(num) //intをstringに変換
	}
}

func main() {
	i := 1
	for i <= 100 {
		fmt.Println(FizzBuzz(i))
		i++
	}
}
