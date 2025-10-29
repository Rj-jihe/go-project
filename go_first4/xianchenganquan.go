// package main
//
// import (
//
//	"fmt"
//	"sync"
//
// )
//
// var sum int
// var wait sync.WaitGroup
// var lock sync.Mutex //加锁
//
//	func add() {
//		lock.Lock() //抢锁
//		for i := 0; i < 100000; i++ {
//			sum += i
//		}
//		lock.Unlock() //开锁
//		wait.Done()
//	}
//
//	func sub() {
//		lock.Lock()
//		for i := 0; i < 100000; i++ {
//			sum -= i
//		}
//		lock.Unlock()
//		wait.Done()
//	}
//
//	func main() {
//		wait.Add(2)
//
//		go add()
//		go sub()
//		wait.Wait()
//		fmt.Println(sum)
//	}
//
// package main
//
// import (
//
//	"fmt"
//	"sync"
//
// )
//
// // 非线程安全的计数器
//
//	type Counter struct {
//		value int
//	}
//
//	func (c *Counter) Increment() {
//		c.value++ // 非原子操作，存在数据竞争
//	}
//
//	func (c *Counter) Value() int {
//		return c.value
//	}
//
//	func main() {
//		counter := &Counter{}
//		var wg sync.WaitGroup
//
//		for i := 0; i < 1000; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				counter.Increment()
//			}()
//		}
//
//		wg.Wait()
//		fmt.Println("Final count:", counter.Value()) // 结果可能小于 1000
//	}
package main

import (
	"fmt"
	"sync"
)

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", counter.Value()) // 总是 1000
}
