package main

import "fmt"

func main() {
	// Quick manual smoke test — feel free to add more
	fmt.Println(ReverseWords("hello world"))             // world hello
	fmt.Println(CaesarEncode("Hello, World!", 3))        // Khoor, Zruog!
	fmt.Println(CaesarDecode("Khoor, Zruog!", 3))        // Hello, World!
	fmt.Println(RLEEncode("aaabbc"))                     // 3a2b1c
	fmt.Println(RLEDecode("3a2b1c"))                     // aaabbc
}
