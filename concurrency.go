
package main

import (
    "fmt"
    "time"
)

// Function representing a task
func worker(id int, results chan string) {
    fmt.Println("Worker", id, "started")

    time.Sleep(4 * time.Second)

    results <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
    // Create a channel
    results := make(chan string)

    // Start measuring execution time
    start := time.Now()

    // Start three goroutines
    go worker(1, results)
    go worker(2, results)
    go worker(3, results)

    // Wait for all workers to finish
    for i := 0; i < 3; i++ {
        message := <-results
        fmt.Println(message)
    }

    // Calculate execution time
    elapsed := time.Since(start)

    fmt.Println("All workers completed!")
    fmt.Println("Total execution time:", elapsed)
}
