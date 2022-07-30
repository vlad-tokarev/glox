package error_reporter

import "fmt"

var (
	HadError bool
)

func Print(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, msg)
	HadError = true
}
