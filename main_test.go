package main

import (
	"testing"
	"encoding/hex"
)

func TestHexGeneration(t *testing.T){
	color := generateHexString("Hello", 6)
	if len(color) != 6 {
		t.Fatal("Size of the generated string must match the passed size")
	}

	_, err := hex.DecodeString(color)
	if err != nil {
		t.Fatalf("Generated string must be a hexadecimal %s", err)
	}
}