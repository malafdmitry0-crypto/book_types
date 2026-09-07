package main

import (
	"fmt"
)

type Job struct{ In, Out []int }

func (j Job) Len() int { return len(j.In) }
func (j Job) Apply(i int) error {
	if j.In[i] < 0 {
		return fmt.Errorf("negative")
	}
	j.Out[i] = 2 * j.In[i]
	return nil
}
func run(j interface {
	Len() int
	Apply(int) error
}) error {
	for i := 0; i < j.Len(); i++ {
		if e := j.Apply(i); e != nil {
			return e
		}
	}
	return nil
}
func main() {
	out := make([]int, 3)
	err := run(Job{[]int{1, -1, 3}, out})
	fmt.Println(out, err != nil)
}
