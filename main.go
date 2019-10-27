package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func ValidatingRows(arr [9][9]rune) bool {
	for k := 0; k < 9; k++ {
		for i := 0; i < 9; i++ {
			for j := i + 1; j < 9; j++ {
				if arr[k][i] != '.' && arr[k][j] != '.' {
					if arr[k][i] == arr[k][j] {
						return false
					}
				}
			}
		}
	}
	return true
}
func ValidatingColumns(arr [9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := j + 1; k < 9; k++ {
				if arr[k][i] != '.' && arr[j][i] != '.' {
					if arr[k][i] == arr[j][i] {
						return false
					}
				}
			}
		}
	}
	return true
}

func ValidatingGrids(arr [9][9]rune) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					for m := i; m < i+3; m++ {
						for n := j; n < j+3; n++ {
							if k != m || l != n {
								if arr[m][n] != '.' {
									if arr[k][l] == arr[m][n] {
										return false
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}

func TotalValid(arr [9][9]rune) bool {
	validRows := ValidatingRows(arr)
	validColumns := ValidatingColumns(arr)
	validGrids := ValidatingGrids(arr)
	if validRows && validGrids && validColumns {
		return true
	} else {
		return false
	}
}

func intToRune(i rune) int {
	if i == '1' {
		return 1
	} else if i == '2' {
		return 2
	} else if i == 3 {
		return '3'
	} else if i == 4 {
		return '4'
	} else if i == 5 {
		return '5'
	} else if i == 6 {
		return '6'
	} else if i == 7 {
		return '7'
	} else if i == 8 {
		return '8'
	} else if i == 9 {
		return '9'
	} else {
		return '0'
	}
}

func SudoSolver(mainArr [9][9]rune, iCord [81]int, jCord [81]int, i int) [9][9]rune {
	if i < 0 {
		return mainArr
	}
	if iCord[i] > 10 {
		return mainArr
	}
	if mainArr[iCord[i]][jCord[i]] < '0' {
		mainArr[iCord[i]][jCord[i]] = '0'
	}
	mainArr[iCord[i]][jCord[i]] = (mainArr[iCord[i]][jCord[i]] + 1)
	if mainArr[iCord[i]][jCord[i]] > '9' {
		mainArr[iCord[i]][jCord[i]] = '.'
		return SudoSolver(mainArr, iCord, jCord, i-1)
	}

	if TotalValid(mainArr) {
		return SudoSolver(mainArr, iCord, jCord, i+1)
	} else {
		return SudoSolver(mainArr, iCord, jCord, i)
	}
}
func main() {
	arguments := os.Args
	var mainArr [9][9]rune
	for index, item := range arguments {
		if index > 0 {
			for innerIndex, itemOfItems := range item {
				mainArr[index-1][innerIndex] = itemOfItems
			}

		}
	}
	var iCord [81]int
	var jCord [81]int
	v := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if mainArr[i][j] == '.' {
				iCord[v] = i
				jCord[v] = j
				v++
			}
		}
	}
	iCord[v] = 100
	v = 0
	if mainArr == SudoSolver(mainArr, iCord, jCord, v) {
		fmt.Println("ERROR")
	} else {
		mainArr = SudoSolver(mainArr, iCord, jCord, v)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				z01.PrintRune(rune(mainArr[i][j]))
				z01.PrintRune(' ')
			}
			z01.PrintRune('\n')
		}
	}

}
