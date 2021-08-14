package main

import "fmt"

func main() {
	fmt.Println(Reverse("George"))
}

func Reverse(str string) string {
	len := len(str)
	reverse := ""
	for i := (len - 1); i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}
