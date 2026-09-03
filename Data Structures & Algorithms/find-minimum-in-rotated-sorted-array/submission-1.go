func findMin(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	l:=0
	r:=len(nums)-1
	right:=r
	reversed:=false
	for (l<=r){
		for nums[r]<nums[l]{
			r--
			reversed=true
		}
		if reversed{
			return nums[r+1]
		}
		if nums[l]<nums[right]{
			return nums[l]
		}
	}
	return -1
}

