package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Listen関数はあるネットワークポート、この場合TCPポートのlocalhost:8000へ入ってくる接続を
	// 待つオブジェクトnet.Listenerを作成する。
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// ListenerのAcceptメソッドは接続要求がくるまで待機し、接続を表すnet.Connオブジェクトを返す
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// net.Connはio.Writerインターフェースを満足するので書き出すことができる
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		// 接続が切れた場合は書き出しに失敗し、errorが返る
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
