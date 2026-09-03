func findMin(nums []int) int {
    res := nums[0]
    l, r := 0, len(nums)-1
	if nums[l]<nums[r]{
		return nums[l]
	}
    for l <= r {
        m := l + (r-l)/2
        if nums[m] >= nums[l] {
			if nums[l]<res{
				res=nums[l]
			}
            l = m + 1
        } else {
			if nums[m]<res{
				res=nums[m]
			}
            r = m - 1
        }
    }
    return res
}
//
