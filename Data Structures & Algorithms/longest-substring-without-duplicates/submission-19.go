func lengthOfLongestSubstring(s string) int {
	l:=0
	r:=1
	if len(s)==0{
		return 0
	}
	seen:=make(map[byte]int)
	seen[s[l]]=l
	max:=1
	///divddfeef 
	for r<len(s){
		if ind, exists:=seen[s[r]]; exists{
			seen=make(map[byte]int)
			l=ind+1
			seen[s[l]]=ind+1
			r=l+1
		}else{
			seen[s[r]]=r
			r++
		}
		if len(seen)>max{
			max=len(seen)
		}
		
	}
	return max
	}
//dvdf

//dvdfe 
