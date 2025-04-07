package main

import "fmt"

func main() {
	var n_gomos int
	var direcao rune //melhor que string

	fmt.Scanf("%d %c", &n_gomos, &direcao)
	gomos := make([][2]int, n_gomos) //2 slices em um so, pra armazenar dois valores inteiros

	for i := 0; i < n_gomos; i++ {
		fmt.Scan(&gomos[i][0], &gomos[i][1]) //gomos[0] int x - gomos[1] int y
	}

	for i := len(gomos) - 1; i > 0; i-- {
		gomos[i][0] = gomos[i-1][0]
		gomos[i][1] = gomos[i-1][1]
	}
	switch direcao {
	//parte do x
	case 'L':
		gomos[0][0] -= 1
	case 'R':
		gomos[0][0] += 1
	case 'D':
		gomos[0][1] += 1
	//parte do y
	case 'U':
		gomos[0][1] -= 1
	}

	for i := 0; i < len(gomos); i++ {
		fmt.Println(gomos[i][0], gomos[i][1])
	}
}

/*for i := 1; i < len(gomos); i++ {
	//guardando numa variavel temporaria o valor  do gomo1
	tX := gomos[i][0]
	tY := gomos[i][1]

	//atualizaa com o gomo anterior
	gomos[i][0] = guardaX
	gomos[i][1] = guardaY

	//atualiza pra proxima posicao
	guardaX = tX
	guardaY = tY
}*/
