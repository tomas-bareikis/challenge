package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	fmt.Println("I will win")
}
