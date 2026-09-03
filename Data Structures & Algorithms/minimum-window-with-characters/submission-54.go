func minWindow(s string, t string) string {
	lenght:=0
    minlenght:=len(s)+1
    og:=make(map[byte]int)
    window:=make(map[byte]int)
    for c:=range t{
        og[t[c]]++
    }
    if len(t)>len(s){
        return ""
    }
    have:=0
    need:=len(og)
    l:=0
    minleft:=0
    for r:=0; r<len(s); r++{
        c:=s[r]
        window[c]++
        lenght++
        if reqtime, exists:=og[c]; exists{
            if window[c]==reqtime{
                have++
            }
        }
        for have==need{
            if _, exists:=og[s[l]]; exists && window[s[l]]==og[s[l]]{
                have--
                if r-l+1<=minlenght {
                    minlenght=lenght
                    minleft=l
                }
            }
            window[s[l]]--
            l++
            lenght--
    
        }
      
    }
    if minlenght > len(s) {
        return ""
    }
    return s[minleft:minleft+minlenght]
}

//xyz z-2 y-5 x-6
///z-2, y-4, x-
// OUZODYXAZVYXYZ"
