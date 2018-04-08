package main

import (
	"strings"
	"path/filepath"
	"fmt"
	"github.com/mhcvs2/godatastructure/util"
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
			if len(suffixes) == 0 {
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
	files, err := util.GetSubFiles("/root")
	suffixes := []string{".txt", ".mp4"}
	if err != nil {
		fmt.Println(err.Error())
	} else {
		out := source(files)
		out2 := filterSuffixes(suffixes, out)
		sink(out2)
	}
}

///root/Untitled Project.mp4
///root/test.txt