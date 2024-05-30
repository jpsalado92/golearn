package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	total_wc := 0
    for _, filename := range os.Args[1:] {
        var line_count, word_count, characer_count int

        file , err := os.Open(filename)

        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            continue
        }

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            line_count++
			total_wc++
			line := scanner.Text()
            word_count += len(line)
            characer_count += len(strings.Fields(line))
        }

		print("line_count: ", line_count, "\n")
		print("word_count: ", word_count, "\n")
		print("characer_count: ", characer_count, "\n")
    }
	print("total_wc: ", total_wc, "\n")
}