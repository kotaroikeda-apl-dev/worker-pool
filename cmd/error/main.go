package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		if job%2 == 0 {
			errors <- fmt.Errorf("Worker %d failed on job %d", id, job)
			continue
		}
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan string, 10)
	errors := make(chan error, 10)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, errors, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)
	close(errors)

	for result := range results {
		fmt.Println(result)
	}
	for err := range errors {
		fmt.Println("Error:", err)
	}
}
