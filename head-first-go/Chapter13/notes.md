# Chapter 13. Sharing work: goroutines and channels
- **Goroutines** let your program work on several different tasks at once.
  - They coordinate their work using **channels**.
  - Which let them send data to each other and synchronize so that oen goroutine doesn't get ahead of another.

## Concurrency using goroutines
- **Concurrency** allows a program to pause on task and work on other tasks.
- If a program is written to support concurrency, then it may also support **parallelism**: running tasks simultaneously.
- In Go, concurrent tasks are called **goroutines**
- To start another goroutine, you use a `go` statement, which is just an ordinary function or method call with the `go` keyword in front of it.

## Go statement can't be used with return values
- Go won't let you use the return value from a function called with a `go` statement, because there is not guarantee the return value will be ready before we attempt to use it.
- There is a way to communicate between goroutines: **channels**
  - Allow you to send values from one goroutine to another
  - They ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it.
- The only practical way to use a channel is to communicate from one goroutine to another goroutine.
- Each channel only carries value of a particular type
  - You might have one channel for `int` values, another channel for values with a struct type.
  - To declare a variable that holds a channel, use the `chan` keyword, followed by the type of values that channel will carry.
    - `var myChannel chan float64`
  - To actually create a channel, you need to call the `make` function.
    - `myChannel = make(chan float64)`
    - or `myChannel := make(chan float64)`

## Sending and receiving values with channels
- To send a value on a channel, use the `<-` operator
  - `myChannel <- 3.14`
- You also use the `<-` to receive values from a channel, but you place the arrow to the left of the channel you are receiving from.

```go
func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
}
```

## Synchronizing goroutines with channels
- Channels do this by **blocking** - by pausing all further operations in the current goroutine.
  - A send operation blocks the sending goroutine until another goroutine executes a receive operation on the same channel.
  - This behavior allows goroutines to **synchronize** their actions

```go
func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	// Create two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Pass each channel to a function running in a new goroutine
	go abc(channel1)
	go def(channel2)

	// Receive and print values from the channels, in order
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
```

## Observing goroutine synchronization
