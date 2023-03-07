package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	From string
	To   string
	// todo: add required flags
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.From, "from", "", "file to read. by default - stdin")
	flag.StringVar(&opts.To, "to", "", "file to write. by default - stdout")

	// todo: parse and validate all flags

	flag.Parse()

	return &opts, nil
}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	fmt.Println(opts)

	// todo: implement the functional requirements described in read.me
}
