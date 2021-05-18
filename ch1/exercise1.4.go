package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise1.4: %v\n", err)
			continue
		} else if checkDuplicate(data) {
			fmt.Printf("duplicate file: %v\n", filename)
		}

	}
}

func checkDuplicate(file []byte) bool {
	counts := make(map[string]int)
	for _, line := range strings.Split(string(file), "\n") {
		counts[line]++
		if counts[line] > 1 {
			return true
		}
	}
	return false
}