package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	pedra_A := make([]int, n) //tipo, tam
	pedra_B := make([]int, n)
	pont := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d%d", &pedra_A[i], &pedra_B[i])
		if pedra_A[i] < 10 || pedra_B[i] < 10 {
			pont[i] = math.MaxFloat64
		} else {
			pont[i] = math.Abs(float64(pedra_A[i] - pedra_B[i]))
		}
	}
	winner := -1
	menor := math.MaxFloat64

	for j := 0; j < n; j++ {
		if pont[j] < menor {
			menor = pont[j]
			winner = j
		}
	}
	if winner == -1 {
		fmt.Println("Sem ganhador")
	} else {
		fmt.Println(winner)
	}
}
