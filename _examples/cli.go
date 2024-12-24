package main

import (
	"fmt"
	"flag"

	"github.com/Kaamkiya/nanoid-go"
)

func main() {
	length := flag.Int("length", nanoid.DefaultLength, "The length of the NanoID to output")
	charset := flag.String("charset", nanoid.DefaultCharset, "The characters the NanoID can use")
	amount := flag.Int("amount", 10, "The amount of NanoIDs to output")

	flag.Parse()

	for i := 0; i < *amount; i++ {
		fmt.Println(nanoid.Nanoid(*length, *charset))
	}
}
