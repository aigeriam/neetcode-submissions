func checkInclusion(s1 string, s2 string) bool {
    og := make(map[byte]int)
    for i := range s1 {
        og[s1[i]]++
    }
    l := 0
    newmap := make(map[byte]int)

    for r := 0; r < len(s2); r++ {
        c := s2[r]
        if _, exists := og[c]; !exists {
            newmap = make(map[byte]int)
            l = r + 1
            continue
        }
        newmap[c]++
        for newmap[c] > og[c] {
            newmap[s2[l]]--
            l++
        }
        if r-l+1 == len(s1) {
            return true
        }
    }
    return false
}

/// lecaabcce
// c a a
/// b-1
/// l=3
///r=4