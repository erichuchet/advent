package file

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	f, err := os.Open(filename)
	exception(err)
	all, err := ioutil.ReadAll(f)
	exception(err)

	return strings.Split(string(all), "\n")
}

func Read(f *io.Reader, err error) *io.Reader {
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func exception(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
