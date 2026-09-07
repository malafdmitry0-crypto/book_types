package main

import (
	"fmt"
)

type Job struct{ In, Out []int }

func (j Job) Apply(i int) { j.Out[i] = j.In[i] * 10 }
func main() {
	s := []int{1, 2, 3}
	j := Job{s[:2], s[1:]}
	for i := 0; i < 2; i++ {
		j.Apply(i)
	}
	fmt.Println(s)
}
