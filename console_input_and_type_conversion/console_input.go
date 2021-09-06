package main

// import multiple mpdules
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //implicitly inffer
	fmt.Printf("Type the year you were born: ")
	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	fmt.Printf("You will be %d years old at the end of 2020  \n", 2020-input)
}
