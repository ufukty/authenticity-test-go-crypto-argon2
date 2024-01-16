# Interoperability of Go Argon2 packages

This repository meant to test compliance of [official Go implementation](https://pkg.go.dev/golang.org/x/crypto/argon2) of Argon2 and Go bindings of [P-H-C-/phc-winner-argon2](https://github.com/P-H-C/phc-winner-argon2) shared by [tvdburgt](<(https://pkg.go.dev/github.com/tvdburgt/go-argon2)>) which is originally written in C.

It compares hashes of top 1000 commonly used passwords produced with both packages with different cost combinations of memory, iterations, threads and length. Than counts the mismatches per combo.

## Dependencies

-   https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt save as `10-million-password.txt`

## Results

```sh
/usr/local/go/bin/go test -timeout 100000s -run ^Test_MostUsedPasswords$ github.com/ufukty/go-argon2-test -v -count=1

=== RUN   Test_MostUsedPasswords
using log file: /var/folders/5l/yb5490cj54zdwxpcr2ypjvzr0000gn/T/3700066352
=== RUN   Test_MostUsedPasswords/mem=1000_threads=1_iter=1_len=32
=== RUN   Test_MostUsedPasswords/mem=2000_threads=1_iter=1_len=32
=== RUN   Test_MostUsedPasswords/mem=1000_threads=2_iter=1_len=32
=== RUN   Test_MostUsedPasswords/mem=1000_threads=1_iter=2_len=32
=== RUN   Test_MostUsedPasswords/mem=1000_threads=1_iter=1_len=64
=== RUN   Test_MostUsedPasswords/mem=2000_threads=2_iter=2_len=64
=== RUN   Test_MostUsedPasswords/mem=4000_threads=2_iter=2_len=64
=== RUN   Test_MostUsedPasswords/mem=2000_threads=4_iter=2_len=64
=== RUN   Test_MostUsedPasswords/mem=2000_threads=2_iter=4_len=64
=== RUN   Test_MostUsedPasswords/mem=2000_threads=2_iter=2_len=128
=== RUN   Test_MostUsedPasswords/mem=4000_threads=4_iter=4_len=128
=== RUN   Test_MostUsedPasswords/mem=8000_threads=4_iter=4_len=128
=== RUN   Test_MostUsedPasswords/mem=4000_threads=8_iter=4_len=128
=== RUN   Test_MostUsedPasswords/mem=4000_threads=4_iter=8_len=128
=== RUN   Test_MostUsedPasswords/mem=4000_threads=4_iter=4_len=256
=== RUN   Test_MostUsedPasswords/mem=8000_threads=8_iter=8_len=256
--- PASS: Test_MostUsedPasswords (105.81s)
    --- PASS: Test_MostUsedPasswords/mem=1000_threads=1_iter=1_len=32 (1.11s)
    --- PASS: Test_MostUsedPasswords/mem=2000_threads=1_iter=1_len=32 (2.41s)
    --- PASS: Test_MostUsedPasswords/mem=1000_threads=2_iter=1_len=32 (1.01s)
    --- PASS: Test_MostUsedPasswords/mem=1000_threads=1_iter=2_len=32 (2.28s)
    --- PASS: Test_MostUsedPasswords/mem=1000_threads=1_iter=1_len=64 (1.19s)
    --- PASS: Test_MostUsedPasswords/mem=2000_threads=2_iter=2_len=64 (3.00s)
    --- PASS: Test_MostUsedPasswords/mem=4000_threads=2_iter=2_len=64 (5.78s)
    --- PASS: Test_MostUsedPasswords/mem=2000_threads=4_iter=2_len=64 (2.24s)
    --- PASS: Test_MostUsedPasswords/mem=2000_threads=2_iter=4_len=64 (5.64s)
    --- PASS: Test_MostUsedPasswords/mem=2000_threads=2_iter=2_len=128 (3.03s)
    --- PASS: Test_MostUsedPasswords/mem=4000_threads=4_iter=4_len=128 (7.84s)
    --- PASS: Test_MostUsedPasswords/mem=8000_threads=4_iter=4_len=128 (15.02s)
    --- PASS: Test_MostUsedPasswords/mem=4000_threads=8_iter=4_len=128 (6.94s)
    --- PASS: Test_MostUsedPasswords/mem=4000_threads=4_iter=8_len=128 (15.51s)
    --- PASS: Test_MostUsedPasswords/mem=4000_threads=4_iter=4_len=256 (7.44s)
    --- PASS: Test_MostUsedPasswords/mem=8000_threads=8_iter=8_len=256 (25.24s)
PASS
ok  	github.com/ufukty/go-argon2-test	106.331s
```

## License

Shared under Apache2 license. See LICENSE file for details.
