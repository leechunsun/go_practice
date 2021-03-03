package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./aa.jaon", os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	a, err := file.Write([]byte("{\"zz\": 13}"))
	fmt.Println(a, err)

}

