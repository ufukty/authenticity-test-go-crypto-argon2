package goargon2test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	tvdburgt "github.com/tvdburgt/go-argon2"
	crypto "golang.org/x/crypto/argon2"
	"golang.org/x/exp/slices"
)

func loadMostUsedPasswords(n int) ([][]byte, error) {
	f, err := os.Open("10-million-password.txt")
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()
	pwds := [][]byte{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pwds = append(pwds, scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning: %w", err)
	}
	return pwds[:min(n, len(pwds))], nil
}

type costCombo struct {
	Memory      int // KiB
	Parallelism int // threads
	Iterations  int // time
	HashLen     int // KeyLen
}

func (c costCombo) String() string {
	return fmt.Sprintf("mem=%d threads=%d iter=%d len=%d", c.Memory, c.Parallelism, c.Iterations, c.HashLen)
}

func Test_MostUsedPasswords(t *testing.T) {
	n := 1000
	pwds, err := loadMostUsedPasswords(n)
	if err != nil {
		t.Fatal(fmt.Errorf("reading most used passwords from file: %w", err))
	}

	tcs := []costCombo{
		{1000, 1, 1, 32},
		{2000, 1, 1, 32},
		{1000, 2, 1, 32},
		{1000, 1, 2, 32},
		{1000, 1, 1, 64},
		{2000, 2, 2, 64},
		{4000, 2, 2, 64},
		{2000, 4, 2, 64},
		{2000, 2, 4, 64},
		{2000, 2, 2, 128},
		{4000, 4, 4, 128},
		{8000, 4, 4, 128},
		{4000, 8, 4, 128},
		{4000, 4, 8, 128},
		{4000, 4, 4, 256},
		{8000, 8, 8, 256},
	}

	l, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		t.Fatal(fmt.Errorf("creating log file in temp: %w", err))
	}
	defer l.Close()
	fmt.Println("using log file:", l.Name())
	var logger = log.New(l, "", 0)

	for _, tc := range tcs {
		t.Run(tc.String(), func(t *testing.T) {
			mismatch := 0

			tvdburgtcontext := &tvdburgt.Context{
				Iterations:  tc.Iterations,
				Memory:      tc.Memory,
				Parallelism: tc.Parallelism,
				HashLen:     tc.HashLen,
				Mode:        tvdburgt.ModeArgon2id,
				Version:     tvdburgt.Version13,
			}

			for i, pwd := range pwds {
				tvdburgthash, err := tvdburgt.Hash(tvdburgtcontext, pwd, []byte("somesalt"))
				if err != nil {
					t.Fatal(fmt.Errorf("hashing %q with tvdburgt: %w", pwd, err))
				}

				cryptoHash := crypto.IDKey(pwd, []byte("somesalt"), uint32(tc.Iterations), uint32(tc.Memory), uint8(tc.Parallelism), uint32(tc.HashLen))

				if slices.Compare(tvdburgthash, cryptoHash) != 0 {
					mismatch += 1
					logger.Printf("assert, mismatch for %q (idx: %d):\n%-10s: %s\n%-10s: %s\n\n", pwd, i, "tvdburgt", tvdburgthash, "crypto", cryptoHash)
				}
			}

			if mismatch != 0 {
				t.Fatalf("mismatched hashes: %d (out of %d)", mismatch, n)
			}

		})
	}
}
