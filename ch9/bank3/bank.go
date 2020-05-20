package bank

import "sync"

var (
	// 慣習的にMutexで保護する変数はMutexの宣言の直後に宣言する
	mu      sync.Mutex // balanceを保護する(排他的ロック)
	balance int
)

func Deposit(amount int) {
	//
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	// あるゴルーチンがLockを獲得していたら、その他のゴルーチンはUnlockされるまで
	// Lock()の呼び出しを待つ
	mu.Lock()
	// Unlock()することで他のゴルーチンがLock()を呼び出せるようになる
	defer mu.Unlock()
	return balance
}
