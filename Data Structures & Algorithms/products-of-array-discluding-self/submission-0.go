func productExceptSelf(nums []int) []int {
	var res[]int
	sum:=1
	for i :=0; i<len(nums); i++{
		sum=1
		k:=i-1
		for k>=0{
			sum=sum*nums[k]
			k--
		}
		j:=i+1
		for j<len(nums){
			sum*=nums[j]
			j++
		}
		res=append(res, sum)
	}
	return res
}///arr[1] arr[2]*arr[3]*arr[4]*arr[5]
