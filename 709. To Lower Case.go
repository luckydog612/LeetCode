package main

func toLowerCase(str string) string {
	var a []byte
	for _, b := range str{
		if b >= 'A' && b <= 'Z'{
			b = b + 32
			a = append(a, byte(b))
		}else{
			a = append(a, byte(b))
		}
	}
	return string(a)
}
