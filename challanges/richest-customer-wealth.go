package challanges

func MaximumWealth(accounts [][]int) int {
	var maxWealth int

	for i := range accounts {
		customerWealth := 0
		for j := range accounts[i] {
			customerWealth += accounts[i][j]
		}

		if customerWealth > maxWealth {
			maxWealth = customerWealth
		}
	}

	return maxWealth
}
