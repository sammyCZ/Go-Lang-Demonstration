package main

import (
    "fmt"
    "time"
)

// Function representing a task
func task(id int, results chan string) {
    fmt.Println("Task", id, "started")

    time.Sleep(4 * time.Second)

    results <- fmt.Sprintf("Task %d finished", id)
}

func main() {
    // Create a channel for go routines
    results := make(chan string)

    // Start measuring execution time
    start := time.Now()

    // Start three goroutines
    go task(1, results)
    go task(2, results)
    go task(3, results)
    
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
