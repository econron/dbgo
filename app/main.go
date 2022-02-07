package main

import (
	"fmt"
	"github.com/econron/dbgo/app/files"
	"sync"
)

func main() {
	fmt.Println("Hello golang from docker!")
	files.Karioki()

	var empty struct{}
	var point struct {
		ID   string
		x, y int
	}
	var array [5]int
	// 配列の値とデータの数は、一度決めたら固定される
	arrayLiteral := [5]int{1, 2, 3, 4, 5}
	arrayInterface := [...]int{1, 2, 3, 4, 5}
	arrayIndex := [...]int{2: 1, 5: 5, 7: 13}

	// データの数は可変。配列のように登録されない
	var slice []int
	sliceLiteral := []int{1, 2, 3, 4, 5}

	// マップの定義
	var m map[string]int
	mapLiteral := map[string]int{
		"John":    42,
		"Ricahrd": 33,
	}
	src := []int{1, 2, 3, 4}
	fmt.Println(src, len(src), cap(src))
	src = append(src, 5)
	fmt.Println(src, len(src), cap(src))
	// 悩んだ時は型のポインタを指定する

	sliceMake := make([]int, 2, 3)
	sliceIndex := []int{2: 1, 3: 5, 7: 13}

	// Slice Tricks
	src2 := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src2))
	copy(dst, src2)
	fmt.Println(dst, len(dst), cap(dst))

	// appendでスライス同士を連結する
	src1, src2 := []int{1, 2}, []int{3, 4, 5}
	dst := append(src1, src2)

	src := []int{1, 2, 3, 4, 5}
	i := 2
	dst := append(src[:i], src[i+1:]...)
	src = []int{1, 2, 3, 4, 5}
	dst = src[:i+copy(src[i:], src[i+1:])]

	// スライスを逆順に並び替える
	src := []int{1, 2, 3, 4, 5}
	for i := len(src)/2 - 1; i >= 0; i-- {
		opp := len(src) - 1 - i
		src[i], src[opp] = src[opp], src[i]
	}

	// スライスの要素を、偶数のみでフィルタリングする
	src := []int{1, 2, 3, 4, 5}
	dst := src[:0]
	for _, v := range src {
		if even(v) {
			dst = append(dst, v)
		}
	}
	fmt.Println(dst)
	for i := len(dst); i < len(src); i++ {
		src[i] = 0
	}
	// マップの初期化
	mapEmpty := map[string]int{}
	mapMake := make(map[string]int)
	mapCap := make(map[string]int, 10)

	// 値およびキーの存在有無の取得
	m := map[string]int{
		"John": 42,
		"Richard": 33,
	}
	age := m["John"]
	fmt.Println(age)

	age, ok := m["Jane"]
	fmt.Println(age, ok)

	_, ok = m["Richard"]
	fmt.Println(ok)

	// 構造体の設計
	// パッケージ外からも参照できる
	type Export struct {
		Name string
		age int
	}

	type unexport struct {
		Name string
		age int
	}

	type Counter struct {
		Name string
		m sync.Mutex
		count int
	}

	func(c *Counter) Increment() int {
		c.m.Lock()
		defer c.m.Unlock()
		c.count++
		return c.count
	}

	func (c *Counter) View() int {
		c.m.RLock()
		defer c.m.RUnlock()
		return c.count
	}

	c := &syncutil.Counter{
		Name: "Access",
	}

	fmt.Println(c.Increment())
	fmt.Println(c.View())
}

func even(n int) bool {
	return n%2 == 0
}