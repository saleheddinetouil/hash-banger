package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
// Import other necessary packages
)

func crackBruteForce(hash, charSet string, minLength, maxLength int, hashType string) {

	for length := minLength; length <= maxLength; length++ {
		generateCombinations(charSet, length, func(combination string) {


			hashed := HashString(combination, hashType)


			if hashed == hash {
				fmt.Println("Password found:", combination)
				os.Exit(0)  // Stop once the password is found
			}
		})
	}
	fmt.Println("Password not found.")
}

// generateCombinations generates all possible combinations of characters from charset up to length
func generateCombinations(charset string, length int, callback func(string)) {

	var wg sync.WaitGroup //using waitgroup for thread safety and efficient concurrency

	ch := make(chan string)


	go func() {
		defer close(ch)
		generate(ch, charset, length, "")
	}()

	for combination := range ch {
		wg.Add(1) // Increment the WaitGroup counter for each combination

        go func(combination string) { // Use a new goroutine for each callback
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			callback(combination)
		}(combination)

	}


	wg.Wait() // Wait for all combinations to be processed


}

// generate is a recursive helper function for bruteforce
func generate(ch chan string, charset string, length int, current string) {
	if length == 0 {
		ch <- current
		return
	}


	for _, char := range charset {
		generate(ch, charset, length-1, current+string(char))
	}


}



func crackDictionary(hash, dictionaryFile string, hashType string) {
	file, err := os.Open(dictionaryFile)
	if err != nil {
		log.Fatal("Error opening dictionary file:", err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()


		hashed := HashString(word, hashType)


		if hashed == hash {
			fmt.Println("Password found:", word)
			return
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading dictionary file:", err)
	}

	fmt.Println("Password not found.")


}
