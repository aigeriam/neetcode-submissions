func twoSum(numbers []int, target int) []int {
	r:=0
	l:=len(numbers)-1

	for r<l{
		if numbers[r]+numbers[l]>target{
			l--
		}else if numbers[r]+numbers[l]<target{
			r++
		}else if numbers[r]+numbers[l]==target{
			return []int{r+1, l+1}
		}		
	}
	return nil 
}

///
///  1 3 5 6 7 8       7
///  1 2 3 5  7 1+6.  7
/// 