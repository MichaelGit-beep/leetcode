package challanges

func NumberOfSteps(num int) int {
	var steps int
	for num >j= 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}
	return num
}
