package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nomuyoshi/go-learning/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)

	// クロール対象のURLをチャネルに送信
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	// チャネルに送信されたクロール対象のURLを1つずつ処理していく
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					// クロールし、新たに見つかったURLをチャネルに送信
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
