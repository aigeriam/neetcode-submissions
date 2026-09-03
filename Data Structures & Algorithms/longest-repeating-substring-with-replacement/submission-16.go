
func characterReplacement(s string, k int) int {
    l := 0
    maxWindow := 0
    maxOccurrence := 0
    occurrenceInWindow := make(map[byte]int)

    for r := 0; r < len(s); r++ {
        c := s[r]
        occurrenceInWindow[c]++
        
        // Track the highest frequency we've seen in the current window configuration
        if occurrenceInWindow[c] > maxOccurrence {
            maxOccurrence = occurrenceInWindow[c]
        }

        // Current window size is (r - l + 1). 
        // If (total characters in window) - (most frequent character) > k,
        // it means we don't have enough replacements left to make the whole window uniform.
        if (r - l + 1) - maxOccurrence > k {
            occurrenceInWindow[s[l]]--
            l++ // Just shrink the window by 1 from the left
        }

        // The maximum valid window size seen so far
        if (r - l + 1) > maxWindow {
            maxWindow = r - l + 1
        }
    }

    return maxWindow
}
///a max windiw will be in case if window-the often elemetn count=<k
/// abcaadddddde.   2

///aaaaa
