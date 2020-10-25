package rwmutex

import (
	"fmt"
	"sync"
	"time"
)

type Test struct {
	lock sync.RWMutex
	items []string
}

func rwmutex() {
	test := Test{}

	go func() {
		for {
			test.lock.Lock()
			fmt.Println("协程1加锁")
			time.Sleep(5 * time.Second)
			test.items = append(test.items, "协程1")
			test.lock.Unlock()
			fmt.Println("协程1解锁")
		}
	}()

	go func() {
		for {
			test.lock.Lock()
			fmt.Println("协程2加锁")
			test.items = append(test.items, "协程2")
			test.lock.Unlock()
			fmt.Println("协程2解锁")
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
	//for {
	//	for _, v := range test.items {
	//		fmt.Println(v)
	//	}
	//}
}