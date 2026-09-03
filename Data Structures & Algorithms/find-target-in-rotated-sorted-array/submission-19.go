func search(nums []int, target int) int {
	l, r:=0, len(nums)-1
	for l<=r{
		m:=l+(r-l)/2
		if nums[m]==target{
			return m
		}
		///increasing
		if nums[l]<=nums[m]{
			/// 3 4 5 6 7 0 1 2
			if target>nums[m] || target<nums[l]{
				l=m+1
			}else{
				r=m-1
				}
		}else{
			if target<nums[m] || target>nums[r]{
				r=m-1
			}else{
				l=m+1
			}
			//it means the start is somwhere between 
		}
	}
	return -1
}





