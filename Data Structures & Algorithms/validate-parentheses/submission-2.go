func isValid(s string) bool {
	var stack[]rune
	pais:=map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
	}
	if len(s)<=1{
		return false
	}
   	for _, r:=range s{
	if r=='[' || r=='(' || r=='{' {
		stack=append(stack, r)
	}else if  (r==']' || r=='}' || r==')') && len(stack)>0{
		if stack[len(stack)-1]==pais[r]{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}
	}else{
		return false 
	}
	}
	if len(stack)==0{
		return true
	}
	return false 

	} 

