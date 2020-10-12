package main

import "fmt"

func main() {
	var value, tmp int

	fmt.Scan(&value)

	s := make([]int, value)

	for i := 0; i < value; i++ {
		fmt.Scan(&s[i])
	}

	for _, v := range s {
		tmp += v
	}

	fmt.Println(tmp)
}
