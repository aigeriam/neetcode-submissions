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
//so like if we are given an array which might be not rotated then its first element will be always lower than the last element, so we first take this situation then we do binary search to see, we take middle potential result and we compare it to most left elemnt of array, if our middle bigger than that it means potetial min could be the most left, and then we find new min and if for that min the left side is bigger, the start of array could be in between so we decrement the right= mid-1 and save potetial result 
/// so basically if we see the array some portion of array is increasing we save the first element, and see if the start of array at the right side. but if we see that the left element is bigger then the middle one could be the start or some elemtn in between the middle one and left elemt 