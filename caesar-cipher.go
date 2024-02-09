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
	return shiftText(message, -key)
}

func shiftText(text string, key int) string {
	var result string
	for _, char := range text {
		if unicode.IsLetter(char) {
			offset := rune('A')
			if unicode.IsLower(char) {
				offset = rune('a')
			}
			adjustedShifted := char - offset + rune(key)
			if adjustedShifted < 0 {
				adjustedShifted = 26 - (adjustedShifted * -1)
			}
			charShifted := adjustedShifted % 26 + offset
			result += string(charShifted)
		}else {
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
		fmt.Println("-e\t -encrypt\t encrypt the message")
		fmt.Println("-d\t -decrypt\t decrypt the message")
		fmt.Println("-k\t -key\t\t key that will be used to encrypt and decrypt the message, if not selected, key 13 will be used as default")
		fmt.Println("-m\t -message\t\t message that will be encrypted")
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