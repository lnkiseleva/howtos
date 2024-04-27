/**
 * Задача:
 * - доработать функцию Do, чтобы она отрабатывала приблизительно за 10мс
 */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TestTaskWb() {
	_, err := Do(context.Background(), []User{{"Paul"}, {"Jack"}, {"Jack"}, {"Mike"}})
	fmt.Print(err)
}

type User struct {
	Name string
}

func fetch(ctx context.Context, u User) (string, error) {

	// что-то делаем по сети
	time.Sleep(time.Millisecond * 10) // имитация задержки

	// if u.Name == "Mike" {
	// 	return "", errors.New("1")
	// }

	return u.Name, nil
}

func Do(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0)
	var wg sync.WaitGroup
	var mu sync.Mutex
	errChan := make(chan error, len(users))
	wg.Add(len(users))
	for _, u := range users {
		go func() {
			defer wg.Done()
			mu.Lock()
			name, err := fetch(ctx, u)
			if err != nil {
				errChan <- err
			}
			names[name] = names[name] + 1
			mu.Unlock()
		}()
	}

	wg.Wait()
	close(errChan)

	for e := range errChan {
		return nil, e
	}

	// debug
	fmt.Println(names)

	return names, nil
}
