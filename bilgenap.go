package main

import "fmt"

func main() {

for i := 2; i < 101; i++ {
	if i%2 == 1 {
		continue
	}
    fmt.Println("bilangan bulat", i)
}
}