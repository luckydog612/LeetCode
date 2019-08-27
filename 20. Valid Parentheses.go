package main

func isValid(s string) bool {
	flag := false
	m := make(map[byte]byte)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'
	sli := make([]byte, 0)
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	for i, b := range s {
		if i == 0 && (s[i] == ')' || s[i] == ']' || s[i] == '}') {
			return false
		}
		if b == '(' || b == '[' || b == '{' {
			sli = append(sli, byte(b))
		} else if b == ')' || b == ']' || b == '}' {
			left := sli[len(sli)-1]
			sli = sli[:len(sli)-1]
			if m[left] == byte(b) {
				flag = true
			} else {
				return false
			}
		}
		if i == len(s)-1 && len(sli) != 0 {
			flag = false
		}
	}
	return flag
}
