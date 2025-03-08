package main

import (
	"fmt"
	"sync"
	"time"
)

const workerCount = 3 // 並列数

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(1 * time.Second) // ジョブ処理
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)

	// Worker を3つ起動
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// ジョブを投入
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs) // チャンネルを閉じる

	wg.Wait()
	fmt.Println("All jobs completed.")
}