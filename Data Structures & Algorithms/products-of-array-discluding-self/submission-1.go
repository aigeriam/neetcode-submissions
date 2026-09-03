func productExceptSelf(nums []int) []int {
	res:=make([]int, len(nums))
	res[0]=1
	prefix:=1
	for i:=1; i<len(nums); i++{
		prefix*=nums[i-1]
		res[i]=prefix
	}
	postfix:=1
	for i:=len(nums)-2; i>=0; i--{
		postfix*=nums[i+1]
		res[i]=res[i]*postfix
	}
	return res
}///res[0]=1 res[1]=[1] res[2]=[1*2] res[3]=[1*2*4] 
///res[0]=2*4*6 res[1]=[1*4*6] res[2]=1*2*
