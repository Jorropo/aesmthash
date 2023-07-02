package main

import (
	"encoding/hex"
	"io"
	"os"

	"github.com/Jorropo/aesmthash"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	h := aesmthash.Hash(input)
	_, err = os.Stdout.WriteString(hex.EncodeToString(h[:]) + "\n")
	if err != nil {
		panic(err)
	}
}
