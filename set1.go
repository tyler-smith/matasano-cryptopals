package main

import (
	"encoding/hex"
	"fmt"
)

// Challenge1 performs Matasano crypto challenge #1
func Challenge1() {
	fmt.Println(hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}

// Challenge2 performs Matasano crypto challenge #2
func Challenge2() {
	a := hexToBytes("1c0111001f010100061a024b53535009181c")
	b := hexToBytes("686974207468652062756c6c277320657965")

	fmt.Println(hex.EncodeToString(calculateReapeatingXor(a, b)))
}

// Challenge3 performs Matasano crypto challenge #3
func Challenge3() {
	secret := hexToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	key, message := crackSingleByteXor(secret)

	fmt.Println("Key is:", string(key))
	fmt.Println("Message is:")
	fmt.Println(string(message))
}

// Challenge4 performs Matasano crypto challenge #4
func Challenge4() {
	challenges := readHexSliceFile("4.txt")

	var winner []byte
	var maxScore float64

	for _, challenge := range challenges {
		for i := 0; i < 255; i++ {
			attempt := calculateSingleByteXor(challenge, byte(i))
			score := englishScore([]byte(attempt))

			if score > maxScore {
				maxScore = score
				winner = attempt
			}

		}
	}

	fmt.Println(string(winner))
}

// Challenge5 performs Matasano crypto challenge #5
func Challenge5() {
	text := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")

	fmt.Println(hex.EncodeToString(calculateReapeatingXor(text, key)))
}

// Challenge6 performs Matasano crypto challenge #6
func Challenge6() {
	secret := readBase64File("6.txt")

	key, message := crackRepeatingKeyXor(secret)

	fmt.Println("Key probably is:", string(key))
	fmt.Println("Message probably is:", string(message))
}

// Challenge7 performs Matasano crypto challenge #7
func Challenge7() {
	key := []byte("YELLOW SUBMARINE")
	secret := readBase64File("7.txt")

	message := decryptAESECB(secret, key)

	fmt.Println(string(message))
}

// Challenge8 performs Matasano crypto challenge #8
func Challenge8() {
	challenges := readHexSliceFile("8.txt")
	blockSize := 16

	for lineNo, challenge := range challenges {
		if isAESECB(challenge, blockSize) {
			fmt.Println("Line", lineNo, "is probably ECB")
			return
		}
	}
}
