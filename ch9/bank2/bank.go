package bank

var (
	sema    = make(chan struct{}, 1) // balanceへのアクセスをホゾするバイナリセマフォ
	balance int
)

func Deposit(amount int) {
	// トークンを獲得。balanceにアクセスするという宣言みたいなもの。
	// semaの容量は1なので、値が送信されると容量がいっぱいになり、他のgoroutineから
	// semaに送信しようとすると待ち状態になる。
	// => 1つのgoroutineのみがbalanceにアクセスすることを保証できる。
	sema <- struct{}{}
	balance = balance + amount
	<-sema // balanceへのアクセス、更新が完了したらトークンを解放（チャネルを受信して空きを作る）
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
