package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	var resultSlice []string
	var letters []string = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q",
		"r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var characters []string = []string{
		"!", "#", "$", "%", "&", "(", ")", "*", "+", ".", "/",
		":", ";", "=", ">", "?", "@", "[", "\\", "]", "^", "`",
		"{", "|", "}", "~", "/"}
	var nums []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	// fmt.Println(os.Args)

	size := flag.Int("s", 0, "Specify the size of the password. This flag is mandatory.")
	num := flag.Bool("n", false, "Use this flag if you want that your password contain a numbers.")
	chars := flag.Bool("c", false, "Use this flag if you want that your password contain a special characters.")
	flag.Parse()

	if *size != 0 {
		if *num && *chars {
			resultSlice = append(letters, nums...)
			resultSlice = append(resultSlice, characters...)
			generatePassword(*size, resultSlice)
			return
		}
		if *num {
			resultSlice = append(letters, nums...)
			generatePassword(*size, resultSlice)
			return
		}
		if *chars {
			resultSlice = append(letters, characters...)
			generatePassword(*size, resultSlice)
			return
		}
		generatePassword(*size, letters)
	}
}

func generatePassword(size int, symb []string) {
	var psw string

	for i := 0; i < size; i++ {
		randomValue := rand.Intn(len(symb))
		psw += symb[randomValue]
	}
	fmt.Println("New password: ", psw)
}
