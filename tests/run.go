package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	log.Println("run tests")
	files, err := os.ReadDir("tests/cases")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !strings.Contains(f.Name(), ".graphql.go") {
			continue
		}
		fmt.Println("Comparing generated sources for " + f.Name())
		expected, err1 := os.ReadFile("tests/cases/" + f.Name())

		if err1 != nil {
			log.Fatal(err1)
		}

		actual, err2 := os.ReadFile("tests/out/cases/" + f.Name())

		if err2 != nil {
			log.Fatal(err2)
		}

		if !bytes.Equal(expected, actual) {
			dmp := diffmatchpatch.New()

			diffs := dmp.DiffMain(string(expected), string(actual), false)

			fmt.Println(dmp.DiffPrettyText(diffs))
			log.Fatal("Test case failed!")
		}
	}
	fmt.Println("All test cases passed!")
}
