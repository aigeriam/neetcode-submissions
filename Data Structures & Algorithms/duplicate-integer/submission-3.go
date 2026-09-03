func hasDuplicate(nums []int) bool {
    seen:=make(map[int] bool)
    for _, r:=range nums{
        if seen[r]{
            return true
        }
        seen[r]=true
    }
    return false

}
