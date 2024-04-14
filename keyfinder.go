package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func LetterFreq(txt string) []int {
	n := 0
	freq := make([]int, 26)
	letter_idx := 65

	for i := 0; i < 26; i++ {
		c := string(i + letter_idx)
		freq[i] = strings.Count(txt, c)
		n += freq[i]
	}
	return freq
}

func IndexOfCoincidence(txt string) float64 {
	txt_len := float64(len(txt) - 1.0)
	let_freq := LetterFreq(txt)
	div := (txt_len * (txt_len - 1)) / 26.0
	idx_coi := 0.0

	var sum float64
	for i := 0; i < 26; i++ {
		sum = sum + float64(let_freq[i]*(let_freq[i]-1))
	}

	if div == 0 {
		return 0.0
	}

	idx_coi = sum / div
	return idx_coi
}

func ProbKeyLen(ciphertext string) []int {
	const avg_lang_ic float64 = 1.73 // Change here to use another language like portuguese (English = 1.73, Portuguese = 1.94)
	
	var prob_key []float64
	var key_len []int

	for i := 1; i <= 32; i++ { // Max key kength size of 32 for now
		avg_ic := 0.0
		ic := 0.0
		for j := 0; j < i; j++ {
			var buffer bytes.Buffer
			for k := 0; k < len(ciphertext)-1; k++ {
				if k%i == j {
					buffer.WriteString(string(ciphertext[k]))
				}
			}
			ic = ic + IndexOfCoincidence(buffer.String())
		}
		avg_ic = ic / float64(i)
		prob_key = append(prob_key, avg_ic)
	}
	for index := range prob_key {
		if math.Abs(prob_key[index]-avg_lang_ic) < 0.20 {
			key_len = append(key_len, index+1)
		}
	}

	return key_len
}

func Transpose(txt string, k_size int) []string {
	var sub_slice []string
	var sub_buf bytes.Buffer
	for index := 0; index < k_size; index++ {
		for i := index; i < len(txt)-1; i = i + k_size {
			sub_buf.WriteString(string(txt[i]))
		}
		sub_slice = append(sub_slice, sub_buf.String())
		sub_buf.Reset()
	}
	return sub_slice
}

func GuessKey(sub_slice []string) string {
	var key string
	// pt_letters_freq := map[string]float64{
	// 	"A": 8.167,
	// 	"B": 1.492,
	// 	"C": 2.782,
	// 	"D": 4.253,
	// 	"E": 12.702,
	// 	"F": 2.228,
	// 	"G": 2.015,
	// 	"H": 6.094,
	// 	"I": 6.966,
	// 	"J": 0.153,
	// 	"K": 0.772,
	// 	"L": 4.025,
	// 	"M": 2.406,
	// 	"N": 6.749,
	// 	"O": 7.507,
	// 	"P": 1.929,
	// 	"Q": 0.095,
	// 	"R": 5.987,
	// 	"S": 6.327,
	// 	"T": 9.056,
	// 	"U": 2.758,
	// 	"V": 0.978,
	// 	"W": 2.361,
	// 	"X": 0.150,
	// 	"Y": 1.974,
	// 	"Z": 0.074,
	// }
	en_letters_freq := map[string]float64{
		"A": 8.167,
		"B": 1.492,
		"C": 2.782,
		"D": 4.253,
		"E": 12.702,
		"F": 2.228,
		"G": 2.015,
		"H": 6.094,
		"I": 6.966,
		"J": 0.153,
		"K": 0.772,
		"L": 4.025,
		"M": 2.406,
		"N": 6.749,
		"O": 7.507,
		"P": 1.929,
		"Q": 0.095,
		"R": 5.987,
		"S": 6.327,
		"T": 9.056,
		"U": 2.758,
		"V": 0.978,
		"W": 2.361,
		"X": 0.150,
		"Y": 1.974,
		"Z": 0.074,
	}

	lang_letters_freq := en_letters_freq // Change here to use another language like portuguese and uncomment the map above

	for _, text := range sub_slice {
		var score []float64

		for _, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			let_score := 0.0

			decrypted_msg := DecryptText(string(letter), text)
			let_freq := LetterFreq(decrypted_msg)
			for i := 0; i < 26; i++ {
				let_score += (float64(let_freq[i]) * lang_letters_freq[string(i+65)])
			}
			score = append(score, let_score)

		}

		bgst := score[0]
		var letter string

		for index, v := range score {
			if v >= bgst {
				bgst = v
				letter = string(index + 65)
			}
		}

		fmt.Println("Score and Letter:", bgst, letter)
		key += letter
	}
	return key
}