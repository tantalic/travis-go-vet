package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// restore os.Stdout
	w.Close()
	os.Stdout = old

	output := <-outC
	if output != "Hello Travis\n" {
		t.Log(output)
		t.Fail()
	}
}
