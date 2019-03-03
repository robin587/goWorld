package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fillUp(id int) chan string {
	content := make(chan string)

	go func() {
		defer close(content)
		sleepNumber := time.Millisecond * time.Duration(10*rand.Float64())
		for i := 0; i < 50; i++ {
			content <- fmt.Sprintf("routine[%v]-%v", id, i)
			time.Sleep(sleepNumber)
		}

	}()

	return content
}

func mux() {
	channelMap := map[int]chan string{}

	for i := 0; i < 5; i++ {
		channelMap[i] = fillUp(i)
	}

	for {
		if len(channelMap) < 1 {
			break
		}

		for id, ch := range channelMap {
			select {
			case v, stillOpen := <-ch:
				if !stillOpen {
					delete(channelMap, id)
				} else {
					color := "\033[34m"
					if (id & 1) != 0 {
						color = "\033[35m"
					}
					fmt.Printf("\033[47mmessage from %s%s\033[00m\n", color, v)

				}
			default:
			}
		}
	}
}

func unmux() {
	channelMap := map[int]chan string{}

	for i := 0; i < 5; i++ {
		channelMap[i] = fillUp(i)
	}

	for {
		if len(channelMap) < 1 {
			break
		}

		for id, ch := range channelMap {
			v, stillOpen := <-ch
			if !stillOpen {
				delete(channelMap, id)
			} else {
				color := "\033[33m"
				if (id & 1) != 0 {
					color = "\033[32m"
				}
				fmt.Printf("\033[47mmessage from %s%s\033[00m\n", color, v)
			}

		}
	}
}

func main() {
	mux()
	fmt.Println("end of multiplexed version, ie non blocking")
	unmux()
}
