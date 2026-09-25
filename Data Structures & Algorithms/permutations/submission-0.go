func permute(nums []int) [][]int {
	res:=[][]int{}
	current:=[]int{}
	var bc func()
	used := make([]bool, len(nums))
	bc=func(){
		if len(current)==len(nums){
			c:=make([]int, len(nums))
			copy(c, current)
			res=append(res, current)
			return 
		}
		for i:=0; i<len(nums); i++{
			if used[i]{
				continue
			}
			used[i]=true
			current=append(current, nums[i])
			bc()
			current = current[:len(current)-1]
        	used[i] = false
		}
	
	}
	bc()
	return res
}

