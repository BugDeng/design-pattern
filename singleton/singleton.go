package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var ins *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
		fmt.Println("singleton init")
	})

	return ins
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstance()
		}()
	}
	wg.Wait()
}
