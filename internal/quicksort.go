package internal

/*

2,10,4,8,5,3

*/

func parition(nums []int, low int, high int) int {

	p := nums[high]

	i := low - 1

	for j := low; j <= high-1; j++ {

		if nums[j] < p {
			i++
			swap(nums, i, j)
		}

	}

	swap(nums, i+1, high)
	return i + 1

}

func quickSrot(nums []int, low int, high int) {

	if low < high {
		p := parition(nums, low, high)

		quickSrot(nums, low, p-1)
		quickSrot(nums, p+1, high)
	}
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
