package main

import "fmt"

func main() {
	//fmt.Printf("Hello %t", false)
	//fmt.Printf("Hello %b \n", 3435) //binary representation of number 3435
	//fmt.Printf("Hello %X \n", 3435)
	//fmt.Printf("Number: %e \n", 2.3736352462356356356536235653632563256536356)
	//fmt.Printf("Number: %f \n", 2.3736)
	//fmt.Printf("Number: %g \n", 2.373635246235635)
	fmt.Printf("Hello %q \n", "Tim")
	fmt.Printf("Hello %9q \n", "Tim")
	fmt.Printf("Hello %-9q is cool \n", "Tim")

	fmt.Printf("Hello %.2f is cool \n", 3.456)

	var out string = fmt.Sprintf("Number: \n \t %07d is cool", 45)
	fmt.Println(out)

}
