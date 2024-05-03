package challanges

func RunningSum(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		sum := nums[0]
		for j := 0; j < i; j++ {
			sum += nums[j+1]
		}
		res = append(res, sum)
	}

	return res
}
