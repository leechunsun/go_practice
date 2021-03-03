package main

import (
	"sort"
)

type MySort []string

func (m MySort) Len() int {
	return len(m)
}

func (m MySort) Less(x ,y int) bool {
	return len(m[x]) > len(m[y])
}

func (m MySort) Swap(x ,y int){
	m[x], m[y] = m[y], m[x]
}

func Solutions(x []string) int {
	c := MySort(x)
	sort.Sort(c)
	for i:=0;i<len(c);i++{
		for j:=i+1;j<len(c);j++{
			a := make(map[string]bool)
			for _, k:= range c[i]{
				a[string(k)] = false
			}
			condition := false
			for _, k:= range c[j]{
				if _, ok := a[string(k)];ok {
					condition = true
				}
			}
			if condition == false{
				return len(c[j]) * len(c[i])
			}
		}
	}
	return 0
}


func main() {
	print(Solutions([]string{"a", "wqerreq", "ewqyeb", "aaaa"}))
}
