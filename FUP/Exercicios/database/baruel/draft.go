package main

import "fmt"

func main() {
	var n_album, n_baruel int
	fmt.Scan(&n_album, &n_baruel)

	figurinhas := make([]int, n_baruel)

	for i := 0; i < n_baruel; i++ {
		fmt.Scan(&figurinhas[i])
	}
	todasFig := make([]int, n_album+1) //slice com todas as figurinhas de baruel

	for j := 0; j < n_baruel; j++ {
		todasFig[figurinhas[j]]++ //ele só vai somar na mesma posição quando o número do valor i for igual ao índice de todasFig
	}
	repetida := false
	espaco := true
	for i := 1; i <= n_album; i++ {
		if todasFig[i] > 1 {
			repetida = true
			for j := 0; j < todasFig[i]-1; j++ {
				if !espaco {
					fmt.Print(" ")
				}
				fmt.Print(i)
				espaco = false
			}
		}
	}

	if !repetida {
		fmt.Printf("N\n")
	} else {
		fmt.Printf("\n")
	}

	faltando := false
    espaco = true

	for i := 1; i <= n_album; i++ {
		if todasFig[i] == 0 {
			faltando = true
			if !espaco {
				fmt.Print(" ")
			}
			fmt.Print(i)
			espaco = false
		}
	}

	if !faltando {
		fmt.Printf("N\n")
	} else {
		fmt.Printf("\n")
	}

}
