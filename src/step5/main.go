// Step5
// Collectionを使ってみましょう
// まず，配列とスライスの違いを理解しましょう．
// その後，スライスの先頭の要素と最後の要素を消去する removeFirstLast(strSlice []string)[]string 関数を定義しましょう．
// 更に，mapを利用してみましょう．
// 人物名をmapのキーとして，その人物の年齢を値として持つmapを定義しましょう．
// そして，mapを受け取って，すべての人物の名前と年齢を連結した文字列の配列を，辞書順で返す getAllNameAndAgeOrder(strSlice map[string]int)[]string を定義しましょう．
// (例) "Taro":12, "Jiro":19, "Saburo":18 のようなmapを受けとったとき，
// 		"Jiro 19", "Saburo 18", "Taro 12" のような文字列の配列を返すようにしてください．
package main

import (
	"fmt"
)

func removeFirstLast(strSlice []string) []string {
	slice_remove := strSlice[1:4]
	return slice_remove
}

// func getAllNameAndAge(strSlice map[string]int) []string {
// 	str := [5]string{"1","2","3","4","5"}
// 	return str
// }

func main() {
	arr := [3]int{0, 1, 2}     //配列
	slice := make([]int, 3, 5) //スライス
	for i := 0; i < 5; i++ {
		if i > 2 { //iの値が3overだったら、領域確保されている容量部に数を代入
			slice = append(slice, i)
		} else { //それ以外であれば普通に数を入れる
			slice[i] = i
		}
	}
	fmt.Println("--ARRAY--")
	fmt.Println(arr)
	fmt.Println("--SLICE--")
	fmt.Println(slice)

	slice_r := [...]string{"0", "1", "2", "3", "4"}
	fmt.Println(removeFirstLast(slice_r[:]))
}
