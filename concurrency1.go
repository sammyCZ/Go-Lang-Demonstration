package main

import (
    "fmt"
    "time"
)

// Function representing a task
func worker(id int, results chan string) {
    fmt.Println("Worker", id, "started")

    time.Sleep(4 * time.Second)

    fmt.Println("Worker", id, "finished")

    results <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
    // Create a buffered channel
    results := make(chan string, 3)

    // Start measuring execution time
    start := time.Now()

    // Execute workers sequentially
    worker(1, results)
    worker(2, results)
    worker(3, results)

    // Receive all results
    for i := 0; i < 3; i++ {
        <-results
    }

    // Calculate execution time
    elapsed := time.Since(start)

    fmt.Println("All workers completed!")
    fmt.Println("Total execution time:", elapsed)
}