// Step9
// Interfaceを利用して，ポリモーフィズムを実装してみましょう．
// 以下のような構造体を定義してください．
// strongMan : 強い人間
//  - money : 所持金(int)
//  - power : 強さ(int)
//  - battle() : 戦闘を行うメソッドです．強い人間は戦闘を行うと所持金が増えます．(強いので)
//  -          : 戦闘を行ったとき，powerを含む文字列を出力してください．
//             : ※所持金の増加量は任意です．intの最大値を超えないようにしてください．
//  - getMoney() : 所持金を取得するメソッドです．
//  - getPower() : 強さを取得するメソッドです．
// otaku :  オタクくん
//  - money : 所持金(int)
//  - anime : オタクくんの好きなアニメ(string)
//  - battle() : 戦闘を行うメソッドです．オタクくんは戦闘を行うと所持金が減ります．(オタクなので)
//  -          : 戦闘を行ったとき，animeを含む文字列を出力してください．
//             : ※所持金の減少量は任意です．intの最小値を下回らないようにしてください．
//  - getMoney() : 所持金を取得するメソッドです．
//  - getAnime() : オタクくんの好きなアニメを取得するメソッドです．

// これらの構造体に対して適切な human インターフェースを実装し，
// humanインターフェースを受け取り，battleメソッドを実行するfightメソッドを実装してください．
// func fight(h human)

package main

import "fmt"

////////////Interface////////////
type human interface {
	battle()
}

func fight(h human) {
	h.battle()
}

////////////strongMan////////////
type strongMan struct {
	money int
	power int
}

func (enemy *strongMan) battle() {
	enemy.money += 100
	fmt.Println("My power is", enemy.getPower())
}

func (enemy *strongMan) getMoney() int {
	return enemy.money
}

func (enemy *strongMan) getPower() int {
	return enemy.power
}

//////////////otaku//////////////
type otaku struct {
	money int
	anime string
}

func (you *otaku) battle() {
	you.money -= 100
	if you.getMoney() <= 0 {
		you.money = 0
	}
	fmt.Println("My favorite anime is ", you.getAnime())
}

func (you *otaku) getMoney() int {
	return you.money
}

func (you *otaku) getAnime() string {
	return you.anime
}

//////////////main////////////////
func main() {
	var enemy strongMan
	var you otaku

	enemy.money = 100
	enemy.power = 200

	you.money = 500
	you.anime = "Dance Dance Revolution"

	// fmt.Println(enemy.getMoney(), enemy.getPower())
	// enemy.battle()
	// fmt.Println(enemy.getMoney(), enemy.getPower())

	// fmt.Println(you.getMoney(), you.getAnime())
	// you.battle()
	// fmt.Println(you.getMoney(), you.getAnime())

	fight(&enemy)
	fight(&you)
}
