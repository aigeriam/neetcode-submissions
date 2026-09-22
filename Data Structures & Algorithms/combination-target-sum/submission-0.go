func combinationSum(nums []int, target int) [][]int {
	res:=[][]int{}
	temp:=0
	vals:=[]int{}
	var dfs func (int)
    dfs=func(i int){
		if temp>target || i>=len(nums){
			i++
			return 
		}
		if temp==target{
			cop := make([]int, len(vals))
            copy(cop, vals )
            res = append(res, cop)
			i++
			return
		}
		vals=append(vals, nums[i])
		temp+=nums[i]
		dfs(i)
		temp=temp-nums[i]
		vals=vals[:len(vals)-1]
		dfs(i+1)
	}
	dfs(0)
	return res
}
