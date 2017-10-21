package main

import (
	"strings"
	"path/filepath"
	"fmt"
)

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0{
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

func main() {
	files := []string{"test.a", "test.b","test.c","test.d"}
	suffixes := []string{".a", ".d"}
	c1 := source(files)
	c2 := filterSuffixes(suffixes, c1)
	sink(c2)
}
