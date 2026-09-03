func twoSum(numbers []int, target int) []int {
	r:=0
	l:=len(numbers)-1
	for r<l{
		for r<l && numbers[r]+numbers[l]!=target{
			l--
		}
		if numbers[r]+numbers[l]==target{
			pair:=[]int{r+1,l+1 }
			return pair 
		}
		r++
		l=len(numbers)-1
	}
	return []int{}

}
/// 1 2 3 4 5 2 1 4 5      7
/// 1 