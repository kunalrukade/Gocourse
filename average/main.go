package main

import "fmt"

func main() {
	student1 := [3]float64{3, 6, 8}
	student2 := [3]float64{6, 8, 3}

	var sum float64

	sum += student1[0] + student1[1] + student1[2]
	sum += student2[0] + student2[1] + student2[2]

	n := len((student1)) + len(student2)

	fmt.Printf("The average of students is %f", sum/float64(n))

}
