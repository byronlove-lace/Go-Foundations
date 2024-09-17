package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// string in go is just a slice of bytes
	const s = "สวัสดี"

	// == len of []byte
	fmt.Println("Len:", len(s))

	// Indexing like this will iterate over the bytes
	for i := 0; i < len(s); i++ {
		// hex values of chars
		fmt.Printf("%x ", s[i])
	}

	// or we can just do this (non-spaced barbarism)
	fmt.Printf("%x\n", s)

	// or this (refined spaced output)
	fmt.Printf("% x\n", s)

	// escapes any trash values and prints the gooduns
	fmt.Printf("%q\n", s)

	// prints only the ascii chars
	fmt.Printf("%+q\n", s)

	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	/*
		Some people think Go strings are always UTF-8, but they are not: only string literals are UTF-8. As we showed in the previous section, string values can contain arbitrary bytes; as we showed in this one, string literals always contain UTF-8 text as long as they have no byte-level escapes.

		But what about the lower case grave-accented letter ‘A’, à? That’s a character, and it’s also a code point (U+00E0), but it has other representations. For example we can use the “combining” grave accent code point, U+0300, and attach it to the lower case letter a, U+0061, to create the same character à. In general, a character may be represented by a number of different sequences of code points, and therefore different sequences of UTF-8 bytes.

		The Go language defines the word rune as an alias for the type int32, so programs can be clear when an integer value represents a code point. Moreover, what you might think of as a character constant is called a rune constant in Go. The type and value of the expression

		A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration. Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes, and the code point is its value. Here’s an example using yet another handy Printf format, %#U, which shows the code point’s Unicode value and its printed representation:I

		Strings are built from bytes so indexing them yields bytes, not characters. A string might not even hold characters. In fact, the definition of “character” is ambiguous and it would be a mistake to try to resolve the ambiguity by defining that strings are made of characters.
	*/

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
