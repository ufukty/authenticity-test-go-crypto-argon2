package goargon2test

import (
	"fmt"
	"os"
	"testing"
)

func Test_MostUsedPasswords(t *testing.T) {

	f, err := os.Open("10-million-password.txt")
	if err != nil {
		t.Fatal(fmt.Errorf("opening file: %w", err))
	}

}
