func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]int)
    l := 0
    max := 0
    //abacdkcmbc
    for r := 0; r < len(s); r++ {
        if ind, exists := seen[s[r]]; exists && ind >= l {
            l = ind + 1
        }
        seen[s[r]] = r
        if r-l+1 > max {
            max = r - l + 1
        }
    }
    return max
}
