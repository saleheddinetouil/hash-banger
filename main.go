package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

type Result struct {
	password string
	found    bool
}

func main() {
	hashType := flag.String("type", "md5", "Hash type (md5, sha1, sha256, bcrypt)")
	hashValue := flag.String("hash", "", "Hash to crack")
	bruteForce := flag.Bool("b", false, "Use brute-force")
	dictionaryAttack := flag.Bool("d", false, "Use dictionary attack")
	charSet := flag.String("chars", "abcdefghijklmnopqrstuvwxyz", "Character set for brute-force")
	minLength := flag.Int("min", 1, "Minimum password length for brute-force")
	maxLength := flag.Int("max", 8, "Maximum password length for brute-force")
	dictionaryFile := flag.String("dict", "dictionary.txt", "Dictionary file for dictionary attack")
	workers := flag.Int("workers", runtime.NumCPU(), "Number of worker goroutines")

	flag.Parse()

	if *hashValue == "" {
		log.Fatal("Hash value is required. Use -hash <hash>")
	}

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt signal. Shutting down...")
		cancel()
	}()

	startTime := time.Now()

	var result Result
	if *bruteForce {
		result = crackBruteForceParallel(ctx, *hashValue, *charSet, *minLength, *maxLength, *hashType, *workers)
	} else if *dictionaryAttack {
		result = crackDictionaryParallel(ctx, *hashValue, *dictionaryFile, *hashType, *workers)
	} else {
		fmt.Println("Please specify a cracking method (-b for brute-force, -d for dictionary).")
		return
	}

	duration := time.Since(startTime)

	if result.found {
		fmt.Printf("\nPassword found: %s\n", result.password)
		fmt.Printf("Time taken: %v\n", duration)
	} else {
		fmt.Println("\nPassword not found")
	}
}

func crackBruteForceParallel(ctx context.Context, hash, charset string, minLen, maxLen int, hashType string, numWorkers int) Result {
	jobs := make(chan string, numWorkers)
	results := make(chan Result, numWorkers)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, jobs, results, hash, hashType)
	}

	// Progress monitoring
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Add progress indicator
				fmt.Print(".")
			}
		}
	}()

	// Generate and send passwords to workers
	go func() {
		generatePasswords(charset, minLen, maxLen, jobs)
		close(jobs)
	}()

	// Wait for result or completion
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process results
	for result := range results {
		if result.found {
			cancel()
			return result
		}
	}

	return Result{found: false}
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan string, results chan<- Result, hash, hashType string) {
	defer wg.Done()
	
	hasher := getHasher(hashType)
	if hasher == nil {
		log.Printf("Unsupported hash type: %s", hashType)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		case password, ok := <-jobs:
			if !ok {
				return
			}
			if hasher.Compare(password, hash) {
				results <- Result{password: password, found: true}
				return
			}
		}
	}
}

// Add similar parallel implementation for crackDictionaryParallel
