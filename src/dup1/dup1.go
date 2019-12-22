//dup1 prints the text of each line that appears more than once in stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //Scan() returns true if line and false if not (effectively a while loop)
		counts[input.Text()]++
		// eq. line := input.Text()
		//     counts[line] = counts[line] + 1
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		} else {
			fmt.Printf("I found nothing")
		}
	}
}
