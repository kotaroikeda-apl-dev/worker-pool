package main

import (
	"fmt"
	"sync"
	"time"
)

const workerCount = 3

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All jobs completed.")
}
