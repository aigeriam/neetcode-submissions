type TimeMap struct {
	m map[string][]pair

}
type pair struct{
	timestamp int
	value string
}


func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int){
	this.m[key]=append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs:=this.m[key] 
	l:=0
	r:=len(pairs)-1
	if _, exists := this.m[key]; !exists {
       return ""
   }
	for l<=r{
		m:=l+(r-l)/2
		if timestamp==pairs[m].timestamp{
			return pairs[m].value
		}
		if timestamp>pairs[m].timestamp{
			if m==len(pairs)-1 || pairs[m+1].timestamp>timestamp{
				return pairs[m].value
			}
			l=m+1
		}else{
			r=m-1
		}
	}
	
	return ""
}
