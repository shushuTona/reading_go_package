package main

import (
	"fmt"
	"sync"
	"time"
)

type Count struct {
	sum int
	m   sync.Mutex
}

func (c *Count) Add() {
	c.m.Lock()
	defer c.m.Unlock()

	c.sum += 1

	fmt.Printf("sum : %d\n", c.sum)
}

func main() {
	count := Count{}
	for i := 1; i <= 10; i++ {
		go count.Add()
	}

	time.Sleep(time.Second * 3)
}
