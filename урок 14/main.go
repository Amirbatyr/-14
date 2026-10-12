
package main

import "fmt"

func main() {
	// Задача 1
	var arr [3][5]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	max := arr[0][0]
	row, col := 0, 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if arr[i][j] > max {
				max = arr[i][j]
				row = i
				col = j
			}
		}
	}

	fmt.Println(row, col)

	// Задача 2 (Снежинка)
	var snow [11][11]string

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			snow[i][j] = "."
		}
	}

	for i := 0; i < 11; i++ {
		snow[5][i] = "*"
		snow[i][5] = "*"
		snow[i][i] = "*"
		snow[i][10-i] = "*"
	}

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Print(snow[i][j], " ")
		}
		fmt.Println()
	}

	// Задача 3 (Шахматная доска)
	var board [8][8]string

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				board[i][j] = "."
			} else {
				board[i][j] = "*"
			}
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}

	// Задача 4 (Поменять строки)
	var matrix1 [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&matrix1[i][j])
		}
	}

	var i, j int
	fmt.Scan(&i, &j)

	matrix1[i], matrix1[j] = matrix1[j], matrix1[i]

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(matrix1[i][j], " ")
		}
		fmt.Println()
	}

	// Задача 5 (Поменять столбцы)
	var matrix2 [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&matrix2[i][j])
		}
	}

	fmt.Scan(&i, &j)

	for k := 0; k < 4; k++ {
		matrix2[k][i], matrix2[k][j] = matrix2[k][j], matrix2[k][i]
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(matrix2[i][j], " ")
		}
		fmt.Println()
	}
}