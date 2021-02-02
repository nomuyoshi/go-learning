package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// ByAge はsort.Interfaceの実装
type ByAge []User

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	// []Userを年齢順にソート
	users := []User{
		{"Sato", 7},
		{"Suzuki", 55},
		{"Takahashi", 24},
		{"Yamada", 75},
	}
	sort.Sort(ByAge(users))
	fmt.Println("ソート後:", users) // ソート後: [{Sato 7} {Takahashi 24} {Suzuki 55} {Yamada 75}]
	sort.Sort(sort.Reverse(ByAge(users)))
	fmt.Println("逆順ソート後:", users) // 逆順ソート後: [{Yamada 75} {Suzuki 55} {Takahashi 24} {Sato 7}]
}
