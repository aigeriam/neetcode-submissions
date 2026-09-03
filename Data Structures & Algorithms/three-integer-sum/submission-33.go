import "slices"
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	triples:=[][]int{}

	for i, n :=range nums{
		if i>0 && n==nums[i-1]{
			continue
		}
		r:=i+1; l:=len(nums)-1
		goal:=-n
		for r<l{
			if nums[r]+nums[l]==goal {
				triples=append(triples, []int{nums[i], nums[r], nums[l]})
				r++
				for nums[r]==nums[r-1] && r<l{
					r++
				}
	
			}else if r<l && nums[r]+nums[l]>goal{
				l--
			}else if nums[r]+nums[l]< goal{
				r++
			}
		}
		if n>=0{
			break
		}
	
	}
	return triples
	}

	//-2 0 0 2 2 
