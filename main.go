package main

import (
	"fmt"
)

func divisibleByThree(input int) bool {
	if (input % 3) == 0 {
		return true
	}
	return false
}

func divisibleByFive(input int) bool {
	if (input % 5) == 0 {
		return true
	}
	return false
}

func main() {
	baseCalc := 0
	sumGroupByThree := 0
	sumGroupByFive := 0
	total := 0

	groupByThree := make([]int, 1)
	groupByFive := make([]int, 1)

	fmt.Println("Informe um número inteiro: ")

	if _, err := fmt.Scanln(&baseCalc); err != nil {
		fmt.Println("Error:", err)
	}

	for i := 0; i < baseCalc; i++ {
		if divisibleByThree(i) {
			groupByThree = append(groupByThree, i)
			sumGroupByThree += i
		}

		if divisibleByFive(i) {
			groupByFive = append(groupByFive, i)
			sumGroupByFive += i
		}
	}

	total = sumGroupByFive + sumGroupByThree

	fmt.Println("\nNúmero divisiveis por 3:", groupByThree)
	fmt.Println("A soma dos numeros divisiveis por 3 é", sumGroupByThree)

	fmt.Println("\nNúmero divisiveis por 5:", groupByFive)
	fmt.Println("A soma dos numeros divisiveis por 5 é", sumGroupByFive)

	fmt.Println("\nA soma Total é", total)

}
