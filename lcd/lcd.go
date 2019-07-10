package lcd

import "fmt"

// AlwaysReturn1 -
func AlwaysReturn1() int {
	return 1
}

var digits = map[string][3]string{
	"0": [3]string{" _ ", "| |", "|_|"},
	"1": [3]string{"   ", "  |", "  |"},
	"2": [3]string{" _ ", " _|", "|_ "},
	"3": [3]string{" _ ", " _|", " _|"},
	"4": [3]string{"   ", "|_|", "  |"},
	"5": [3]string{" _ ", "|_ ", " _|"},
	"6": [3]string{" _ ", "|_ ", "|_|"},
	"7": [3]string{" _ ", "  |", "  |"},
	"8": [3]string{" _ ", "|_|", "|_|"},
	"9": [3]string{" _ ", "|_|", " _|"},
	".": [3]string{"   ", "   ", "  ."},
}

func LCD(number float32) [3]string {

	str := fmt.Sprint(number)

	var line1 = ""
	var line2 = ""
	var line3 = ""

	for _, ch := range str {
		var char = string(ch)
		// fmt.Println(char)

		var arrDigit = digits[char]
		line1 += arrDigit[0]
		line2 += arrDigit[1]
		line3 += arrDigit[2]
	}
	return [3]string{line1, line2, line3}

	// var result [3]string
	// switch number {
	// case 1:
	// 	result = [3]string{"   ", "  |", "  |"}
	// case 2:
	// 	result = [3]string{" _ ", " _|", "|_ "}
	// case 3:
	// 	result = [3]string{" _ ", " _|", " _|"}
	// case 4:
	// 	result = [3]string{"   ", "|_|", "  |"}
	// case 5:
	// 	result = [3]string{" _ ", "|_ ", " _|"}
	// case 6:
	// 	result = [3]string{" _ ", "|_ ", "|_|"}
	// case 7:
	// 	result = [3]string{" _ ", "  |", "  |"}
	// case 8:
	// 	result = [3]string{" _ ", "|_|", "|_|"}
	// case 9:
	// 	result = [3]string{" _ ", "|_|", " _|"}
	// case 0:
	// 	result = [3]string{" _ ", "| |", "|_|"}
	// }

	// //result[0] = "   "
	// //result[1] = "  |"
	// //result[2] = "  |"
	// return result
}
