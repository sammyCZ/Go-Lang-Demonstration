
package main

import (
    "fmt"
    "time"
)

// Function representing a task
func task(id int) {
    fmt.Println("Task", id, "started")

    time.Sleep(4 * time.Second)

    fmt.Println("Task", id, "finished")
}

func main() {
    // Start measuring execution time
    start := time.Now()

    // Execute tasks sequentially
    task(1)
    task(2)
    task(3)

    // Calculate execution time
    elapsed := time.Since(start)

    fmt.Println("All tasks completed!")
    fmt.Println("Total execution time:", elapsed)
}
