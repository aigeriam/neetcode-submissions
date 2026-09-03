func longestConsecutive(nums []int) int {
    numSet:=make(map[int]struct{})
    for _, num:=range nums{
        numSet[num]=struct{}{}
    }
    longest:=0
   
    for num:=range numSet{
        if _, found:=numSet[num-1]; !found{
            lenght:=1
            for{
                if _, exists:=numSet[num+lenght]; exists{
                    lenght++
                }else{
                    break
                }
            }
            if lenght>longest{
                longest=lenght
            }
        }

        }
        return longest
    }


/// 2,  20