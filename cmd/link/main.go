package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gophercises/link"
)

var htmlFileName *string = new(string)

func init() {
	flag.StringVar(htmlFileName, "file", "", "the name of the html file to parse")
	flag.Parse()
	if *htmlFileName == "" {
		htmlFileName = nil
	}
}

func exit(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	var rdr io.Reader
	var err error = nil

	if htmlFileName == nil {
		rdr = os.Stdin
	} else {
		rdr, err = os.Open(*htmlFileName)
	}

	if err != nil {
		exit(err)
	}

	links, err := link.Parse(rdr)

	if err != nil {
		exit(err)
	}

	for _, l := range links {
		fmt.Printf("href: \"%s\" text: \"%s\"\n", l.Href, l.Text)
	}
}
