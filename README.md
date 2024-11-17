# hash-banger: Password Hash Cracker (Educational Purposes Only)

This tool is designed for educational purposes only to demonstrate various password cracking techniques. It should only be used on hashes you own or have explicit permission to crack.  Unethical use of this tool is strongly discouraged.

## Features

* Brute-force cracking (configurable character set and length).
* Dictionary attack (using a provided wordlist).
* [Future] Rainbow table support.
* Supports multiple hash algorithms (MD5, SHA1, SHA256, bcrypt).

## Getting Started

1. **Build:** `go build`
2. **Run:**  `./hash-banger -h` to see usage options.

## Usage Examples

* **Brute-force:** `./hash-banger -b -hash <hash> -chars <charset> -min <min_len> -max <max_len>`
* **Dictionary attack:** `./hash-banger -d -hash <hash> -dict <dictionary_file>`

## Disclaimer

Use this tool responsibly and ethically.  Cracking passwords without authorization is illegal and harmful.

## Code Structure

* `main.go`:  Command-line argument parsing and program entry point.
* `cracker.go`:  Core cracking logic.
* `utils.go`: Utility functions (hashing, character set generation).


## Contributing

Contributions are welcome! (Especially for rainbow table support)



## License

MIT
