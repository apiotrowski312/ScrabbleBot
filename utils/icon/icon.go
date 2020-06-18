package icon

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func LoadIcon(filename string) []byte {
	iconFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(iconFile)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return b
}
