package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

// mutex.Lock()
// doSomeThing()
// mutex.Unlock()
// 注意，平时所说的锁定，其实就是去锁定互斥锁，而不是说去锁定一段代码。
// 也就是说，当代码执行到有锁的地方时，它获取不到互斥锁的锁定，
// 会阻塞在那里，从而达到控制同步的目的。// need to confirm

// go run -race main.go
// vs
// go run main.go
