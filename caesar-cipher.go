package main

import (
	"flag"
	"fmt"
	"unicode"
)

func caesarEncrypt(message string, key int) string {
	return shiftText(message, key)
}

func caesarDecrypt(message string, key int) string {
	return originalText(message, key)
}

func shiftText(text string, key int) string {
	var result string
	for _, char := range text {
		if unicode.IsLetter(char) {
			offset := rune('A')
			if unicode.IsLower(char) {
				offset = rune('a')
			}
			charShifted := (char - offset + rune(key)) % 26 + offset
			result += string(charShifted)
		}else {
			result += string(char)
		}
	}
	return result
}

func originalText(text string, key int) string{
	var result string
	for _, char := range text{
		if unicode.IsLetter(char){
			offset := rune('A')
			if unicode.IsLower(char){
				offset = rune('a')
			}
			x := char - offset - rune(key)
			if x < 0{
				x = 26 - x + offset
				result += string(x)
			}else {
				charShifted := x % 26 + offset
				result += string(charShifted)
			}
		}else{
			result += string(char)
		}	
	}
	return result
}	

func main() {
	// Definiar as flags de entrada
	encryptFlag := flag.Bool("e", false, "Encrypt a message")
	decryptFlag := flag.Bool("d", false , "Decrypt a message")
	messageFlag := flag.String("m", "", "Message to encrypt or decrypt")
	keyFlag := flag.Int("k", 13, "Key for Caesar Chipher (default is 13)")
	helpFlag := flag.Bool("h", false, "Man Ceaser Cipher")

	flag.Parse()

	// tratar entradas inesperas (ex: comando errado, as opções são: etc)
	
	if *messageFlag == "" || *helpFlag == true{
		fmt.Println("Using: go run caesar-cipher.go [option]")
		fmt.Println("-e, -encrypt     encrypt the message")
		fmt.Println("-d, -decrypt     decrypt the message")
		fmt.Println("-k,  -key        key that will be used to encrypt and decrypt the message, if not selected, key 13 will be used as default")
		return
	}
	if *encryptFlag {
		fmt.Println("Encrypting message...")
		encryptMessage := caesarEncrypt(*messageFlag, *keyFlag)
		fmt.Println(encryptMessage)
	}
	if *decryptFlag {
		fmt.Println("Decrypting message...")
		decryptMessage := caesarDecrypt(*messageFlag, *keyFlag)
		fmt.Println(decryptMessage)
	}
}