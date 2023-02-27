# Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine

- Create new channels
- Send a value into a channel
- Receive a value from the channel
- Deadlock

  ```go
  package main

  func main() {
      ch := make(chan int)
      ch <- 5
  }
  ```

- Unidirectional channels
- Closing channels
- For range loops on channels
- Channel buffering
