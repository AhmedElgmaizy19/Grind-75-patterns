package TwoPointer

import "unicode"


func isValidPlaindroem(s string) bool {
	txt := []rune(s)
	start := 0
	end := len(txt)-1

	for start < end {
		if !unicode.IsLetter(txt[start])&& !unicode.IsNumber(txt[start]){
			start++
			continue
		}
		if !unicode.IsLetter(txt[end]) && !unicode.IsNumber(txt[end]){
			end--
			continue
		}
		if unicode.ToLower(txt[start]) != unicode.ToLower(txt[end]) {
			return false
		}
		start++
		end--

	}
	return true

}



