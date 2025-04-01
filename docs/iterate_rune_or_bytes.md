When iterating over a string in Go, you can do so as a `Rune` or as a `Byte`.

- Byte: an unsigned 8-bit integer (uint8). In Go, bytes are used to represent ASCII characters (character that fits in one byte)
- Rune: 4 bytes = 32 bits, an int32. In Go, runes are used to represent Unicode characters (characters from any language or symbol, requires more than on byte).

```go
	// Indexing a string yields bytes
	fmt.Println("Bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v = %T, ", s[i], s[i])
	}

	fmt.Println("Runes")
	// Ranging a string yields Runes
	for r := range s {
		fmt.Printf("%v = %T, ", r, r)
	}
```
