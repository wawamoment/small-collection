package main
import "fmt"

/*
13.06.2026 - 17.14
  Second ever project. I originally did a very similar project in Python when I was in 3rd grade.. now I'm gonna be 10th grade after
  the summer.. time flies.
  I do plan on integrating the American system.
*/

var res float64
var resr float64

func Turkish(ex, exS float64, project float64) float64 {
	// FIRST EXAM
	if project>=0 {
		res = ex+exS+project
		resr = res/3
	} else {
		res = ex+exS
		resr = res/2
	}
	return resr
}

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Print("Enter your first exams result, and your sözlü note, separated by the space key.\n>>>")
	fmt.Scan(&b, &c)
	fmt.Print("Did you do a project?\nIf yes, please type the points you got. If not, type -1 and it will not be calculated.\n>>>")
	fmt.Scan(&a)

	fmt.Print(Turkish(b,c,a))
}
