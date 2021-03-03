package main

import "fmt"

func deal(origin []int, deg int) int{
	ob_array := make([]int, deg+1)
	for i := 1 ; i <= deg; i++{
		for _, c := range origin{
			if i-c >= 0 && i - c <= deg + 1 {
				if i-c == 0{
					ob_array[i] = 1
					continue
				}
				if ob_array[i-c] > 0 && ob_array[i-c] + 1 > ob_array[1]{
					ob_array[i] = ob_array[i-c] + 1
				}
			}
		}
	}
	fmt.Println(ob_array)
	return ob_array[deg]
}



func main() {
	a := []int{1,3,5}
	deg := 20
	fmt.Println(deal(a, deg))
}