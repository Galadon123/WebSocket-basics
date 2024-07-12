## Goroutine

A goroutine is a lightweight thread of execution managed by the Go runtime. It allows functions to run concurrently with other functions. Here's a basic example to illustrate goroutines:

```go
package main

import (
    "fmt"
    "time"
)

func sayHello(id int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Hello from goroutine %d\n", id)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    // Start two goroutines
    go sayHello(1)
    go sayHello(2)

    // Wait for goroutines to finish
    time.Sleep(time.Second * 3)
    fmt.Println("Main function finished")
}
```

Let's break this down:

1. We define a `sayHello` function that prints a message 5 times with a small delay between each print.

2. In the `main` function, we start two goroutines using the `go` keyword:
   ```go
   go sayHello(1)
   go sayHello(2)
   ```
   This launches two concurrent executions of `sayHello`.

3. We add a sleep in the main function to give the goroutines time to execute before the program exits.

When you run this program, you'll see output similar to this:

```
Hello from goroutine 1
Hello from goroutine 2
Hello from goroutine 1
Hello from goroutine 2
Hello from goroutine 1
Hello from goroutine 2
Hello from goroutine 1
Hello from goroutine 2
Hello from goroutine 1
Hello from goroutine 2
Main function finished
```

Key points about goroutines:

1. Concurrency: Goroutines run concurrently. In this example, both `sayHello` functions are executing at the same time.

2. Lightweight: You can easily create thousands of goroutines without significantly impacting system resources.

3. Communication: Goroutines can communicate with each other using channels (not shown in this example).

4. Scheduling: The Go runtime schedules goroutines onto OS threads, handling the complexities of thread management for you.

5. Asynchronous: Goroutines allow you to perform operations asynchronously, improving program efficiency.
