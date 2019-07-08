package lcd

// AlwaysReturn1 -
func AlwaysReturn1() int {
	return 1;
}

func LCD(number int) [3]string {
	var result [3]string
	switch number {
	case 1: result = [3]string {"   ", "  |", "  |"}
	case 2: result = [3]string {" _ ", " _|", "|_ "}
	case 3: result = [3]string {" _ ", " _|", " _|"}
	case 4: result = [3]string {"   ", "|_|", "  |"}
	case 5: result = [3]string {" _ ", "|_ ", " _|"}
	case 6: result = [3]string {" _ ", "|_ ", "|_|"}
	case 7: result = [3]string {" _ ", "  |", "  |"}
	}
	
	//result[0] = "   "
	//result[1] = "  |"
	//result[2] = "  |"
	return result
}