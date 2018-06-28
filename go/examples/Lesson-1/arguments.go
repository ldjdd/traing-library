package main

import "fmt"
import "os"
import "strings"

func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	sep = ""
	s1 := ""
	for _, arg := range os.Args[1:] {
		s1 += sep + arg
		sep = " "
	}
	fmt.Println(s1)

	fmt.Println(strings.Join(os.Args[1:], " "))
}