func findDuplicate(nums []int) int {
	hash:=make(map[int]bool)
	for _, n:=range nums{
		if hash[n]{
			return n
		}
		hash[n]=true

	}
	return 1
}
