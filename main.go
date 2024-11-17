package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	hashType := flag.String("type", "md5", "Hash type (md5, sha1, sha256, bcrypt)")
	hashValue := flag.String("hash", "", "Hash to crack")
	bruteForce := flag.Bool("b", false, "Use brute-force")
	dictionaryAttack := flag.Bool("d", false, "Use dictionary attack")
	charSet := flag.String("chars", "abcdefghijklmnopqrstuvwxyz", "Character set for brute-force")
	minLength := flag.Int("min", 1, "Minimum password length for brute-force")
	maxLength := flag.Int("max", 8, "Maximum password length for brute-force")
	dictionaryFile := flag.String("dict", "dictionary.txt", "Dictionary file for dictionary attack")


	flag.Parse()

	if *hashValue == "" {
		log.Fatal("Hash value is required. Use -hash <hash>")
	}


	if *bruteForce {
		crackBruteForce(*hashValue, *charSet, *minLength, *maxLength, *hashType)
	} else if *dictionaryAttack {
		crackDictionary(*hashValue, *dictionaryFile, *hashType)
	} else {

		fmt.Println("Please specify a cracking method (-b for brute-force, -d for dictionary).")

	}
}
