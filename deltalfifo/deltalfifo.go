package deltalfifo

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
	"time"
)

func DelTalFIFO() {
	fifo := cache.NewFIFO(func (obj interface{}) (string, error) {
		o := obj.(map[string]string)
		return o["id"], nil
	})

	i := 0
	go func() {
		for {
			time.Sleep(2*time.Second)
			i += 1
			fifo.Add(map[string]string{
				"id": fmt.Sprintf("%d", i),
				"name": "andrew",
			})
		}
	}()

	//fifo.Add(map[string]string{
	//	"id": "1",
	//	"name": "andrew",
	//})
	//fifo.Add(map[string]string{
	//	"id": "2",
	//	"name": "david",
	//})
	//fifo.Add(map[string]string{
	//	"id": "3",
	//	"name": "jame",
	//})
	//fifo.Add(map[string]string{
	//	"id": "4",
	//	"name": "jack",
	//})

	//fmt.Println(fifo.List())
	//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))
	//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))
	//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))
	//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))
	//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))

	for {
		item, exist, _ := fifo.GetByKey("5")
		if exist {
			fmt.Println("I had find out: ", item)
			fifo.Pop(func(interface{}) error {return nil})
		}
		//fmt.Println(fifo.Pop(func(interface{}) error {return nil}))
	}

}
