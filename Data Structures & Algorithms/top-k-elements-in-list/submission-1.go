func topKFrequent(nums []int, k int) []int {
	res:=make(map[int] int)
	for _, n:=range nums{
		res[n]++
	}
	freq:=make([][]int, len(nums)+1)
	for num, cnt:=range res{
		freq[cnt]=append(freq[cnt], num) 
	}
	final:=[]int{}
	for i:=len(freq)-1; i>0; i--{
		for _, num:=range freq[i]{
			final=append(final, num)
			if len(final)==k{
				return final
			}
		}
		
	}
	return final 
}
