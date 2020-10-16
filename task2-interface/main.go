package main

import (
	"log"
)

type cetak interface {
	rataKiri() string
	rataKanan() string
}


type triangle struct {
	number int
}

type reverseTriangle struct {
	numberB int
}

func (t triangle) rataKiri() string {
		for i := 0; i < t.number; i++ {
		star := "*"
		for j := 0; j < i; j++ {
			star += "*"
		}
		log.Printf(star)
	}
	return ""
}

func (r reverseTriangle) rataKanan() string {
	for i := 0; i < r.numberB; i++ {
		test := ""
		for j := 0; j < i; j++ {
			test += " "
		}
		for k := r.numberB - 1; k >= i; k-- {
			test += "*"
		}
		log.Printf(test)
	}
	return ""
}


func main() {
	var rata cetak
	rata = triangle{5}
	rata.rataKiri()
	
	rata = reverseTriangle{5}
	rata.rataKanan()
}