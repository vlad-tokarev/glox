package main

import (
	"bufio"
	"fmt"
	"github.com/vlad-tokarev/glox/error_reporter"
	"github.com/vlad-tokarev/glox/scanner"
	"io"
	"log"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		runPrompt()
	case 2:
		runFile(os.Args[1])
	default:
		log.Fatalf("Usage: %s <file>", os.Args[0])
	}
}

func runFile(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file: %s", err)
	}
	run(string(dat))
	if error_reporter.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		var line string
		var err error
		if line, err = reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Could not read line: %s", err)
		}
		run(line)
		error_reporter.HadError = false
	}
}

func run(s string) {
	sc := scanner.NewScanner(s)
	for {
		token, err := sc.Next()
		switch {
		case err == scanner.ErrDone:
			return
		case err != nil:
			log.Fatalf("Could not scan: %s", err)
		default:
			fmt.Printf("Token: %#v\n", token)
		}
	}
}
