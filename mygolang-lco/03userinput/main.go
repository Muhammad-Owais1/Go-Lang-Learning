package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wellcome := "hi"
	fmt.Println(wellcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("thank", input)
}
