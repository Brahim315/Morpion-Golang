package main

import (
	"fmt"
	"os"
)

func check(table1 []string, table2 []string, table3 []string) bool {
	if (table1[0] == table3[0] && table2[0] == table3[0]) ||
		(table1[1] == table3[1] && table2[1] == table3[1]) ||
		(table1[2] == table3[2] && table2[2] == table3[2]) ||
		(table1[0] == table3[2] && table2[1] == table3[2]) ||
		(table1[2] == table3[0] && table2[1] == table3[0]) ||
		(table1[0] == table1[2] && table1[1] == table1[2]) ||
		(table2[0] == table2[2] && table2[1] == table2[2]) ||
		(table3[0] == table3[2] && table3[1] == table3[2]) {
		fmt.Println("Game finished !")
		return true
	}
	return false
}

func main() {
	var count int
	var table1 = []string{"1", "2", "3"}
	var table2 = []string{"4", "5", "6"}
	var table3 = []string{"7", "8", "9"}

	fmt.Println("Start a game ? Y/N")
	var resp string
	fmt.Scanln(&resp)
	if resp == "Y" || resp == "y" {
		fmt.Println("Player 1, choose a name : ")
		var player1 string
		fmt.Scanln(&player1)
		fmt.Println("Player 2, choose a name : ")
		var player2 string
		fmt.Scanln(&player2)
		fmt.Print(table1, "\n", table2, "\n", table3, "\n")
		for check(table1, table2, table3) != true {
			fmt.Println(player1 + " : Your turn, select a number between 1 and 9")
			var tour1 string
			fmt.Scanln(&tour1)
			for i := 0; i < len(table1); i++ {
				for j := 0; j < len(table2); j++ {
					for k := 0; k < len(table3); k++ {
						if tour1 == table1[i] {
							table1[i] = "X"
						}
						if tour1 == table2[j] {
							table2[j] = "X"
						}
						if tour1 == table3[k] {
							table3[k] = "X"
						}
					}
				}
			}
			count++
			if count == 9 {
				fmt.Println("No winner...")
				main()
			}
			fmt.Print(table1, "\n", table2, "\n", table3, "\n")
			if check(table1, table2, table3) == true {
				fmt.Println(player1 + " won !")
				main()
			}
			fmt.Println(player2 + " : Your turn, select a number between 1 and 9")
			var tour2 string
			fmt.Scanln(&tour2)
			for i := 0; i < len(table1); i++ {
				for j := 0; j < len(table2); j++ {
					for k := 0; k < len(table3); k++ {
						if tour2 == table1[i] {
							table1[i] = "O"
						}
						if tour2 == table2[j] {
							table2[j] = "O"
						}
						if tour2 == table3[k] {
							table3[k] = "O"
						}
					}
				}
			}
			count++
			fmt.Println(count)
			if count == 9 {
				fmt.Println("No winner...")
				main()
			}
			fmt.Print(table1, "\n", table2, "\n", table3, "\n")
			if check(table1, table2, table3) == true {
				fmt.Println(player2 + " won !")
				main()
			}
		}
	}
	if resp == "N" || resp == "n" {
		fmt.Println("OK, bye")
		os.Exit(3)
	}
	fmt.Println("Please select Y or N")
	main()
}
