func twoSum(nums []int, target int) []int {
	stored:=make(map[int] int)
	for i, n:=range nums{
		dif:=target-n
		if f, ok:=stored[dif]; ok{
			return []int{f, i}
		}
		stored[n]=i
}
return nil 
}
