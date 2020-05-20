package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nomuyoshi/go-learning/ch5/links"
)

// tokens は並列数を20個に制限するためのchannel
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 実行前にtokensに送信。容量がいっぱいの場合は、空きがでるまで待機する
	list, err := links.Extract(url)
	<-tokens // 実行後にtokensから受信、値を解放することでchannelの空きを増やす
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	var n int // worklistへの送信待ち数をカウントする

	// クロール対象のURLをチャネルに送信
	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	// チャネルに送信されたクロール対象のURLを1つずつ処理していく
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				// 未完了のURLがあれば、nをカウントアップし、crawlを実行
				n++
				go func(link string) {
					// クロールし、新たに見つかったURLをチャネルに送信
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
