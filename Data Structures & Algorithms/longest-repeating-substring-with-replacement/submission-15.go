func characterReplacement(s string, k int) int {
	r:=0
	l:=0
	max:=0
	occuranceinwindow:=make(map[byte]int)
	maxoccurance:=0
	themostfrequentelment:=s[0]
	for r<len(s){
		window:=r-l+1
		c:=s[r]
		occuranceinwindow[c]++
		if occuranceinwindow[c]>=maxoccurance{
			maxoccurance=occuranceinwindow[c]
			themostfrequentelment=c
		}
		for window-maxoccurance>k{
			if themostfrequentelment==s[l]{
				maxoccurance--
				themostfrequentelment, maxoccurance=findthenextfrequency(occuranceinwindow)
			}
			occuranceinwindow[s[l]]--
			l++
			window--
		}
		if window>max{
			max=window
		}
		r++
	}
	return max
}
func findthenextfrequency(m map[byte]int)(byte, int){
	frequent:=0
	var el byte
	for k, v:=range m{
		if v>frequent{
			frequent=v
			el=k
		}
	}
	return el, frequent
}
///a max windiw will be in case if window-the often elemetn count=<k
/// abcaadddddde.   2
