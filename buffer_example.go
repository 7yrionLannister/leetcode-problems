package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer        // A Buffer needs no initialization.
	b.Write([]byte("Hello ")) // write to buffer
	fmt.Fprintf(&b, "world!\n")
	b.WriteTo(os.Stdout)
	b.ReadFrom(os.Stdin)
	fmt.Println("Your input to the buffer is", b.String())
}
