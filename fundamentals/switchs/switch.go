package main

func weeklyDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	default:
		return "undefined"
	}
}

func main() {
	weeklyDay(2)
}
