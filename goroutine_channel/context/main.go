package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// context.Background() は空のContextを生成する
	// WithCancelでキャンセルできるようにする
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err := printGreeting(ctx); err != nil {
			fmt.Printf("cannot print greeting: %v\n", err)
			// エラーが返ってきたらmainからContextをキャンセルする
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// printGreeting がエラーを返して Context がキャンセルされたらことは
		// printFarewell にも ctx を通して伝搬する
		if err := printFarewell(ctx); err != nil {
			fmt.Printf("cannot print farewell: %v\n", err)
		}
	}()

	wg.Wait()
}

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	// err は Context がキャンセルされたときのエラー
	if err != nil {
		return err
	}

	fmt.Printf("%s world!\n", greeting)
	return nil
}

func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(ctx context.Context) (string, error) {
	// 1sec でタイムアウトする子Contextを生成
	// 1sec 後に Context はキャンセルされる
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Context がキャンセルされたことをlocaleに伝搬するためにctxを渡す
	// err は Context がキャンセルされたときのエラー
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "goodbye", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	// err は Context がキャンセルされたときのエラー
	switch locale, err := locale(ctx); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}

	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		// キャンセルされたら、エラーを返す。
		// このエラーはmainまで伝搬する
		return "", ctx.Err()
	case <-time.After(10 * time.Second):
	}

	return "EN/US", nil
}
