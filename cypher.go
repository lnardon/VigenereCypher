package main

func IsLetter(char byte) bool {
        if char >= 65 && char <= 90 {
                return true
        }
        return false
}

func DecryptChar (textChar byte, keyChar byte) string {
        dc := (textChar - keyChar) % 26

        for dc < 0 {
                dc += 26
        }

        return string(dc + 65)
}

func DecryptText(txt string, key string) string {
	res := ""
	k_idx := 0
	c := 0
	for i := 0; i < len(txt); i++ {
                if IsLetter(txt[i]) {
                        k_idx = c % len(key)
                        res += DecryptChar(txt[i], key[k_idx])
                        c++
                }
	}

	return res
}
