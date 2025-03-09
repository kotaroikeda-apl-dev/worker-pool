package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d exiting\n", id)
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	// Worker 数を動的に変える
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	time.Sleep(5 * time.Second) // 負荷がかかったと想定して Worker を増やす
	wg.Add(1)
	go worker(ctx, 4, jobs, &wg)

	close(jobs)
	wg.Wait()
	fmt.Println("All jobs completed.")
}
