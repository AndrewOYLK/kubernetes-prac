package synccond

import (
	"fmt"
	"sync"
	"time"
)

func synccond () {
	str := ""
	cond := sync.Cond{}

	go func() {
		for {
			if str == "" {
				cond.Wait()
			}
			fmt.Println("Find:", str)
		}
	}()

	for {
		time.Sleep(2 * time.Second)
		if str == "" {
			str = "andrew"
		} else {
			str = ""
		}
	}
}
