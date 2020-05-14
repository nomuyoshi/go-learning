package main

import "fmt"

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0 // reset
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12
}

// ByteCounter はバイト数をカウントする
// インタフェースio.Writerを満たすようにWriteメソッドを実装する
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}
