func twoSum(nums []int, target int) []int {
    sp:=make(map[int]int)
    for ind, i :=range nums{
        diff:=target-i
        if j, found :=sp[diff]; found{
            return []int{j, ind}
        }
        sp[i]=ind


    }
    return []int{}
    
}