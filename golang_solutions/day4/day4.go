package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var secretKey string
	if len(os.Args) < 2 {
		secretKey = "yzbqklnj"
	} else {
		secretKey = os.Args[1]
	}

	fmt.Printf("Answer for the secret key \"%s\" (hash start with 5 zeros) is %d\n", secretKey, FindAnswerForSecretKey(secretKey, "00000"))
	fmt.Printf("Answer for the secret key \"%s\" (hash start with 6 zeros) is %d\n", secretKey, FindAnswerForSecretKey(secretKey, "000000"))
}

func FindAnswerForSecretKey(s, hashStartWith string) int {
	var i int = 1
	for {
		input := PrepareHashInput(s, i)
		hash := CalculateHash(input)
		if IsHashStartsWith(hash, hashStartWith) {
			return i
		}
		i++
	}
}

func PrepareHashInput(s string, i int) []byte {
	return []byte(s + strconv.Itoa(i))
}

func IsHashStartsWith(hash, with string) bool {
	return hash[:len(with)] == with
}

func CalculateHash(in []byte) string {
	return fmt.Sprintf("%x", md5.Sum(in))
}
