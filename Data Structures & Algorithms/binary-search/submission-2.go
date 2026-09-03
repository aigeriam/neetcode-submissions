func search(nums []int, target int) int {
	n:=len(nums)/2
	l:=0
	r:=len(nums)-1
	for l<=r{
		if nums[n]==target{
			return n
		}else if nums[n]>target{
			r=n-1
		}else if nums[n]<target{
			l=n+1
		}
		n=(l+r)/2
	}
	return -1

}



