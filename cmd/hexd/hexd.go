package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/hexd"
)

func readBytes() []byte {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		return data
	}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	fmt.Println(hexd.DumpString(readBytes()))
}
