package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    fields := strings.Fields(s)
    counts := make(map[string]int, len(fields))
    for _, word := range fields {
        counts[word] += 1
    }
    return counts
}

func main() {
    wc.Test(WordCount)
}