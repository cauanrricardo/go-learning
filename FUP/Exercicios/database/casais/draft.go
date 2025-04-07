package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	var cont = 0
	vetor := make([]int, n) //tipo - tamanho

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	for j := 0; j < n; j++ {
		for k := 0; k < n; k++ {
			if vetor[j] == -vetor[k] && vetor[k] != 0 { //o j so trocar o valor quando verificar todos os elementos de k
				cont++
				vetor[j] = 0
				vetor[k] = 0
				break
			}
		}
	}
	fmt.Println(cont)

}
