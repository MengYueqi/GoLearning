package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var cond = sync.NewCond(&mutex)
var dataReady = false

// 消费者
func consumer(id int) {
	cond.L.Lock()
	for !dataReady { // 等待条件满足
		fmt.Printf("Consumer %d is waiting\n", id)
		cond.Wait() // 进入等待并释放锁
	}
	fmt.Printf("Consumer %d consumed the data\n", id)
	cond.L.Unlock()
}

// 生产者
func producer() {
	time.Sleep(2 * time.Second) // 模拟生产数据
	cond.L.Lock()
	dataReady = true
	fmt.Println("Producer produced the data")
	cond.Signal()
	cond.L.Unlock()
}

func main() {
	go consumer(1)
	go consumer(2)
	go producer()

	time.Sleep(5 * time.Second) // 等待所有 goroutine 执行完毕
}
