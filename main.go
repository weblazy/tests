package main

import (
	"fmt"

	"github.com/weblazy/tests/aes"
)

func main() {
	keys := "sfdasfd"
	tokenBytes, err := aes.NewAes([]byte(keys)).Encrypt("hah1")
	token := string(tokenBytes)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", token)
	res, err := aes.NewAes([]byte(keys)).Decrypt("DXlSfoh+kyRRDXgLc2n3gQ==")
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", res)

}
