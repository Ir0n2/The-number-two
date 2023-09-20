package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceRoll(num int) int {

   	s1 := rand.NewSource(time.Now().UnixNano())
    	r1 := rand.New(s1)
	rand := r1.Intn(num)
    	return rand
}

func main() {
var d int
	for i := 0; i <= 100; i++ {

		r := diceRoll(6)
		fmt.Println(r)
		if r == 2 {
			d++
			fmt.Println("here")
		} else {
			continue
		}
	}
                fmt.Println("The number two occurs ", d, "times")
}

