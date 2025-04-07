package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 //criei as variaveis
	fmt.Scan(&a, &b, &c) //fiz a leitura com scan

	var p float64 = (a + b + c) / 2 //calc o p
	var area float64 = math.Sqrt(p * (p - a) * (p - b) * (p - c)) //area ao ^2
	fmt.Printf("%.2f\n", area) //printar na formt correto (como no c)

}
