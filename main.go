package main

import (
	"os"
	"time"

	"example.com/learnGoWithTests/mocking"
)

//
// func main() {
// 	list := linkedList.LinkedList{}
// 	list.Add(1)
// 	list.Add(2)
// 	list.Add(3)
// 	list.Add(4)
// 	list.Show()
// 	value, _ := list.Remove(2)
// 	fmt.Println(value)
// 	list.Show()
// 	list.Add(0)
// 	list.Show()
// }

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
