package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/rohrschacht/mailmanmembers"
)

func main() {
	in := flag.String("in", "", "Input file containing html")
	out := flag.String("out", "", "Output file")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	if *in != "" {
		f, err := os.Open(*in)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open file %s: %v\n", *in, err)
			os.Exit(2)
		}

		defer f.Close()

		reader = bufio.NewReader(f)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()

	members, err := mailmanmembers.MembersFromString(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error parsing the html: %v\n", err)
		os.Exit(1)
	}

	outputfunc := func(str string) { fmt.Println(str) }

	if *out != "" {
		outf, err := os.Create(*out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not write to file %s: %v\n", *out, err)
			os.Exit(3)
		}

		defer outf.Close()

		writer := bufio.NewWriter(outf)
		outputfunc = func(str string) { writer.WriteString(fmt.Sprintf("%s\n", str)) }
		defer writer.Flush()
	}

	for _, member := range members {
		outputfunc(member)
	}
}
