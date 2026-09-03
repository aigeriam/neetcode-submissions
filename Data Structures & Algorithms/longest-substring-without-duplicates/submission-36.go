func lengthOfLongestSubstring(s string) int {
	l:=0
	r:=1
	if len(s)==0{
		return 0
	}
	//cabcbdkac
	seen:=make(map[byte]int)
	seen[s[l]]=l
	max:=1
	///pwwkew w
	for r<len(s) {
		if ind, exists:=seen[s[r]]; exists{
			if ind>=l{
				l=ind+1
			}
			
		}
		seen[s[r]]=r
		if r-l+1>max{
			max=r-l+1
		}
		r++
		
	}
	return max
	}
//dvdf

//dvdfe 
