package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
    "golang.org/x/crypto/bcrypt"

)



func HashString(password, hashType string) string {
	switch hashType {

	case "md5":
		hasher := md5.New()

		hasher.Write([]byte(password))
		return hex.EncodeToString(hasher.Sum(nil))



	case "sha1":
		hasher := sha1.New()
		hasher.Write([]byte(password))
		return hex.EncodeToString(hasher.Sum(nil))



	case "sha256":

		hasher := sha256.New()

		hasher.Write([]byte(password))
		return hex.EncodeToString(hasher.Sum(nil))



	case "bcrypt":


        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {

            fmt.Println("Error hashing password with bcrypt:", err) // Handle the error as needed


            return ""

        }


        return string(hashedPassword)


	default:
		panic("Unsupported hash type") // Or handle it more gracefully.
	}
}
