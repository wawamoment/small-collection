package main
import "fmt"

// first Go project! -- 12.06.2026 -- 21.04 V1 -|- 13.06.2026 -- 15.32 -- V2

func add(n1, n2 int) int {
	addi := n1+n2
	return addi
}
func subs(n1, n2 int) int {
	subst := n1-n2
	return subst
}
func mult(n1, n2 int) int {
	multi := n1*n2
	return multi
}
func divi(n1, n2 int) int {
	divid := n1/n2
	return divid
}

func main() {
	var choice int
	var n1 int
	var n2 int
	fmt.Printf("\nSelect your choice:\n1 - Add\n2 - Substract\n3 - Multiply\n4 - Divide\n>>> ")
	fmt.Scan(&choice)
	if choice == 1 {
		fmt.Print("Please type the numbers you want to add (space in between):\n>>> ")
		fmt.Scan(&n1, &n2)
		fmt.Print(add(n1,n2))
	} else if choice == 2 {
		fmt.Print("Please type the numbers you want to substract (space in between):\n>>> ")
		fmt.Scan(&n1, &n2)
		fmt.Print(subs(n1,n2))
	} else if choice == 3 {
		fmt.Print("Please type the numbers you want to multiply (space in between):\n>>> ")
		fmt.Scan(&n1, &n2)
		fmt.Print(mult(n1,n2))
	} else if choice == 4 {
		fmt.Print("Please type the numbers you want to divide (space in between):\n>>> ")
		fmt.Scan(&n1, &n2)
		fmt.Print(divi(n1,n2))
	}
}
