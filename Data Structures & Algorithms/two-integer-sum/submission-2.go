func twoSum(nums []int, target int) []int {
	stored:=make(map[int] int)
	for i, n:=range nums{
		stored[n]=i
	}
	for i:=0; i<len(nums); i++{
	    val, exists:=stored[target-nums[i]]
		if exists && val!=i{
			return []int{i, stored[target-nums[i]]}
		}
	}
	return nil
}

