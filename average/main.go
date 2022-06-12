package main

import "fmt"

func main() {
	students := [2][3]float64{
		{8, 5, 4},
		{7, 8, 6},
	}

	var sum float64

	sum += students[0][0] + students[0][1] + students[0][2]
	sum += students[1][0] + students[1][1] + students[1][2]

	n := len(students) * len(students[0])

	fmt.Printf("The average of students is %f", sum/float64(n))

}
