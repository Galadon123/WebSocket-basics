
## Buffered and unbuffered Channels in Go:

Channels are used for communication between goroutines. There are two types: buffered and unbuffered.

1. Unbuffered Channels:
   - No capacity to hold data
   - Sender blocks until receiver is ready

Example:
```go
func main() {
    ch := make(chan int) // Unbuffered channel
    
    go func() {
        ch <- 1 // This will block until someone receives
        fmt.Println("Sent")
    }()
    
    time.Sleep(time.Second) // Simulate work
    fmt.Println(<-ch) // Receive
    time.Sleep(time.Second) // Wait to see "Sent"
}
```
Output:
```
1
Sent
```
Explanation: The goroutine blocks on sending until the main goroutine receives.

2. Buffered Channels:
   - Can hold a limited number of values
   - Sender only blocks when buffer is full

Example:
```go
func main() {
    ch := make(chan int, 2) // Buffered channel with capacity 2
    
    ch <- 1 // These don't block
    ch <- 2
    fmt.Println("Sent two numbers")
    
    fmt.Println(<-ch) // 1
    fmt.Println(<-ch) // 2
}
```
Output:
```
Sent two numbers
1
2
```
Explanation: Sending doesn't block because the channel has buffer space. Receiving works as normal.

Key Difference:
- Unbuffered: Synchronizes sender and receiver
- Buffered: Allows sending without immediate reception, up to buffer capacity

Choose based on your synchronization needs: use unbuffered for strict synchronization, buffered for more flexibility.