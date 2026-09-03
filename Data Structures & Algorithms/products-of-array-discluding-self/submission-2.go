func productExceptSelf(nums []int) []int {
	arr:=make([]int, len(nums))
	///1 2 4 6.   [0]=1 [1]=1 [2]=1*2  [3]=1*2*4 
	/// [3]= post=6 
	prev:=1
	for i:=0; i<=len(nums)-1; i++{
		arr[i]=prev
		prev*=nums[i]
	}
	post:=1
	for j:=len(nums)-1; j>=0; j--{
		arr[j]*=post
		post*=nums[j]
	}
	return arr

}
