package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func main() {
	secret := "yzbqklnj"
	i := 0
	for {
		h := md5.New()
		newSecret := fmt.Sprintf("%s%d", secret, i)
		io.WriteString(h, newSecret)
		hash := h.Sum(nil)
		hexHash := hex.EncodeToString(hash)
		if strings.HasPrefix(hexHash, "000000") {
			fmt.Println(newSecret)
			fmt.Println(hexHash)
			break
		}
		i++
	}
}
