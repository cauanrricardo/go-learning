package main

import "fmt"

func main() {
	var h, p, f, d int32
	fmt.Scan(&h, &p, &f, &d)

	fugiu := true

	if d == 1 {
		for i := f; i != h; i = (i + 1) % 16 {
			if i == p {
				fugiu = false
				break
			}
		}
	} else {
		for i := f; i != h; i = (i - 1 + 16) % 16 { //+ 16 pra n dar negativo
			if i == p {
				fugiu = false
				break
			}
		}
	}

	if fugiu {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
	// sentindo horario - diminui
	// sentindo anti-horario - aumenta
}
