package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

type Result struct {
	Words int
	Lines int
	Bytes int
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	countLines := flag.Bool("l", false, "Count lines")

	// Defining a boolean flag -b to additionally count Bytes
	countBytes := flag.Bool("b", false, "Count Bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	res, err := count(os.Stdin, *countLines, *countBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	if *countLines {
		fmt.Printf("Number of lines: %d\n", res.Lines)
	} else {
		fmt.Printf("Number of Words: %d\n", res.Words)
	}

	if *countBytes {
		fmt.Printf("Number of Bytes: %d", res.Bytes)
	}

	fmt.Println()
}

func count(r io.Reader, countLines bool, countBytes bool) (Result, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return Result{}, err
	}

	res := Result{}

	if countBytes {
		res.Bytes = len(data)
	}

	// Re-scan the in-memory data for either Words or Lines
	sc := bufio.NewScanner(bytes.NewReader(data))

	if countLines {
		// The default Split function is to scan lines so we go straight to counting
		for sc.Scan() {
			res.Lines++
		}
	} else {
		sc.Split(bufio.ScanWords)
		for sc.Scan() {
			res.Words++
		}
	}

	if err := sc.Err(); err != nil {
		return Result{}, err
	}

	return res, nil
}
