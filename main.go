package main

import "fmt"
import "github.com/module/lcd"

func main() {
	result := lcd.LCD(1234567.892)
	fmt.Println(result[0])
	fmt.Println(result[1])
	fmt.Println(result[2])
}

