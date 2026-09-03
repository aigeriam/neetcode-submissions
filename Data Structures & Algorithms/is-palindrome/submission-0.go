func isPalindrome(s string) bool {
	l, r:=0, len(s)-1
	for l<r{
		for l<r && isAlphaNum(rune(s[l])) && isAlphaNum(rune(s[r])){
			if strings.ToLower(string(s[l]))!=strings.ToLower(string(s[r])){
				return false
			}
			l++
			r--
		}
		if !(isAlphaNum(rune(s[l]))){
			l++
		}
		if !(isAlphaNum(rune(s[r]))){
			r--
		}
	}
	return true
	
}
func isAlphaNum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}