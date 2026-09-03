func productExceptSelf(nums []int) []int {
	prod:=make([]int, len(nums))
	sum:=nums[0]
	prod[0]=1
	i:=1
	for i<len(nums){
		prod[i]=sum
		sum*=nums[i]
		i++
	}
	sum=1
	for i=len(nums)-1; i>=0; i--{
		prod[i]*=sum
		sum*=nums[i]
	} 
	return prod
}
