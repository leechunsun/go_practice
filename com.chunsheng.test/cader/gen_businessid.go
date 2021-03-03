package main

import (
	"crypto/md5"
	"fmt"
)

func GenBusinessID(uid int, business string) string {
	resp := fmt.Sprintf("%d:%s", uid, business)
	md := md5.New()
	md.Write([]byte(resp))
	a := md.Sum(nil)
	return fmt.Sprintf("%x", string(a))
}

func main() {
	z := GenBusinessID(125, "repay")
	fmt.Println(z)
}


