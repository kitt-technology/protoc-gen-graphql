package main

import (
	"bytes"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	log.Println("run tests")
	files, err := ioutil.ReadDir("tests/cases")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".graphql.go") {
			fmt.Println("Comparing generated sources for " + f.Name())
			expected, err1 := ioutil.ReadFile("tests/cases/" + f.Name())

			if err1 != nil {
				log.Fatal(err1)
			}

			actual, err2 := ioutil.ReadFile("tests/out/cases/" + f.Name())

			if err2 != nil {
				log.Fatal(err2)
			}

			// Ignore the import statements, as the order is inconsistent

			re := regexp.MustCompile(`(?s)import \([\s\S]*?\)`)

			expectedStr := string(expected)
			expectedStrNoImports := re.ReplaceAllString(expectedStr, "")

			actualStr := string(actual)
			actualStrNoImports := re.ReplaceAllString(actualStr, "")

			expected = []byte(expectedStrNoImports)
			actual = []byte(actualStrNoImports)

			if !bytes.Equal(expected, actual) {
				dmp := diffmatchpatch.New()

				diffs := dmp.DiffMain(string(expected), string(actual), false)
				fmt.Println(dmp.DiffPrettyText(diffs))
				log.Fatal("Test case failed!")
			}
		}
	}
	fmt.Println("All test cases passed!")
}
