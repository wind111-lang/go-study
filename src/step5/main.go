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
	"sort"
)

func removeFirstLast(strSlice []string) []string {
	slice_remove := strSlice[1:4] //1番目から3番目までをのこす[n:m-1]
	//この場合は[1:4-1]なので、実際には[1:3]
	return slice_remove
}

func getAllNameAndAge(strSlice map[string]int) []string {
	//キーのスライスを作成する
	var keys []string
	for key := range strSlice {
		keys = append(keys, key) //配列に要素を追加
	}
	//キーをソートする
	sort.Strings(keys)

	//ソートしたスライスをforで回して繰り返しの順番を固定する
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i], strSlice[keys[i]])
	}

	return keys
}

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
	fmt.Println(arr) //配列
	fmt.Println("--SLICE--")
	fmt.Println(slice) //スライス

	slice_r := [...]string{"0", "1", "2", "3", "4"}
	fmt.Println(removeFirstLast(slice_r[:])) //先頭と末尾を取り除く

	namelist := map[string]int{
		"Taro":   12,
		"Jiro":   19,
		"Saburo": 18,
	}

	getAllNameAndAge(namelist)
}
