package pvi

import (
	"bufio"
	"fmt"
	"os"
)

// PrintToTerminal prints the given string to the terminal
func PrintToTerminal(output string) {
	fmt.Println(output)
}

// PrintToFile prints the given string to the specified file
func PrintToFile(output string, filename string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(output)
	check(err)

	w.Flush()
	f.Sync()
}
