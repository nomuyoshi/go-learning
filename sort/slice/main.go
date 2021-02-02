package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// []stringを文字数（len）が少ない順にソート
	languages := []string{"Ruby", "JavaScript", "Go", "PHP", "Python"}
	sort.Slice(languages, func(i, j int) bool { return len(languages[i]) < len(languages[j]) })
	fmt.Printf("ソート後: %s\n", languages) // ソート後: [Go PHP Ruby Python JavaScript]

	// []Userを年齢順にソート
	users := []User{
		{"Sato", 7},
		{"Suzuki", 55},
		{"Takahashi", 24},
		{"Yamada", 75},
	}
	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	fmt.Println("ソート後:", users) // [{Sato 7} {Takahashi 24} {Suzuki 55} {Yamada 75}]
}
