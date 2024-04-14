package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter ciphered text:")
	reader := bufio.NewReader(os.Stdin)
	cipherText, _ := reader.ReadString('\n')

	fmt.Println("Ciphered Text:", cipherText)

	probableKeyLengths := ProbKeyLen(cipherText)
	fmt.Println("Probable Key Lengths:", probableKeyLengths)

	keyLength := probableKeyLengths[0]
	fmt.Println("Using Key Length:", keyLength)

	transposedText := Transpose(cipherText, keyLength)

	guessedKey := GuessKey(transposedText)
	fmt.Println("Guessed Key:", guessedKey)

	decryptedText := DecryptText(cipherText, guessedKey)
	fmt.Println("Decrypted Text:", decryptedText)
}
