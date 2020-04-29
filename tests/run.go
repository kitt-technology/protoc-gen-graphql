package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main()  {
	log.Println("run tests")
	files, err := ioutil.ReadDir("tests/cases")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".auth.go") {
			fmt.Println("Comparing generated sources for " + f.Name())
			expected, err1 := ioutil.ReadFile("tests/cases/" + f.Name())

			if err1 != nil {
				log.Fatal(err1)
			}

			actual, err2 := ioutil.ReadFile("tests/out/cases/" + f.Name())

			if err2 != nil {
				log.Fatal(err2)
			}

			if !bytes.Equal(expected, actual) {
				log.Fatal("Test case failed!")
			}
		}
	}
	fmt.Println("All test cases passed!")
}
