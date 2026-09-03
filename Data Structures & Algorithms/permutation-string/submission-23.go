func checkInclusion(s1 string, s2 string) bool {
    og := make(map[byte]int)
    for i := range s1 {
        og[s1[i]]++
    }
    l := 0
    r:=0
    newmap := make(map[byte]int)
    for r< len(s2) {
        c:=s2[r]
        occuranceinog, exists:=og[c]
        if  !exists{
            newmap=make(map[byte]int) ///left is start 
            r++
            l=r
            continue
        }
        newmap[c]++
        for newmap[c]>occuranceinog{
            newmap[s2[l]]--
            l++
        }
        if r-l+1==len(s1){
            return true
        }
        r++
}
return false 
}
/// lecaabcce
// c a a
/// b-1
/// l=3
///r=4