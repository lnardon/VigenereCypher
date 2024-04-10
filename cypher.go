package main

func IsLetter(char byte){
        if char >= 65 && char <= 90 {
                return true
        }
        return false
}

func DecryptChar (textChar byte, keyChar byte) {
        dc := (textChar - keyChar) % 26

        for dc < 0 {
                dc += 26
        }

        return string(dc + 65)
}

func EncryptChar (textChar byte, keyChar byte){
        ec := ((textChar - 65) + (keyChar - 65)) % 26
        return string(ec + 65)
}

func DecryptText(text string, key string){
        result := ""
}

func EncryptText(text string, key string){
        result := ""
}
