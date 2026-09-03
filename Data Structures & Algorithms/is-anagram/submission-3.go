
func isAnagram(s string, t string) bool {
	seen:=make(map[byte] int)
	seen2:=make(map[byte] int)
	if len(s)!=len(t){
		return false
	}
	for i:=0; i<len(s); i++{
		seen[s[i]]+=1
		seen2[t[i]]+=1
	}
	for key :=range seen{
		if seen[key]!=seen2[key] || len(s)!=len(t){
			return false
		}
	}
	return true 
	
}
