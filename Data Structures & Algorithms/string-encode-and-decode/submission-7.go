type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	str:=""
	for _, s:=range strs{
		str+=strconv.Itoa(len(s))+"#"+s
	}
	return str
}

func (s *Solution) Decode(encoded string) []string {
	res:=[]string{}
	i:=0;
	for i<len(encoded){
		j:=i
		for encoded[j]!='#'{
			j++
		}
		lenght,_:=strconv.Atoi(encoded[i:j])
		i=j+1
		res=append(res, encoded[i:i+lenght])
		i+=lenght

	}
	return res
}
