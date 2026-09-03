func topKFrequent(nums []int, k int) []int {
	mymap:=make(map[int] int)
	for _, i:=range nums{
		mymap[i]++
	}
	sorted := make([][]int, len(nums)+1)
	for num, count:=range mymap{
		sorted[count]=append(sorted[count], num)
	}
	//'0, 0, 0, 7, {9}, 0, 0, 0'
	var res[]int
	for i:=len(sorted)-1; i>=0 && k>0;  i--{
		for j:=len(sorted[i])-1; j>=0 && k>0; j--{
				res=append(res, sorted[i][j])
				k--
			}
		}
		return res
	}
	

