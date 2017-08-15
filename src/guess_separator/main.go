package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"log"
	"bufio"
	"io"
	"errors"
)

func main()  {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s file\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	separators := []string{"\t", "*", "|", "."}
	linesRead, lines := readUpToLines(os.Args[1], 5)
	counts := createCount(lines, separators, linesRead)
	separator := guessSep(counts, separators, linesRead)
	report(separator)
}

func readUpToLines(file string, number int) (linesRead int, lines []string) {
	inFile := os.Stdin
	err := errors.New("")
	if inFile, err = os.Open(file); err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	eof := false
	linesRead = 0
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		lines = append(lines, line)
		linesRead++
		if err == io.EOF || linesRead >= number {
			err = nil
			eof = true
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return linesRead, lines
}

func createCount(lines, separators []string, linesRead int) [][]int {
	counts := make([][]int, len(separators))
	for sepIndex := range separators {
		counts[sepIndex] = make([]int, linesRead)
		for lineIndex, line := range lines {
			counts[sepIndex][lineIndex] = strings.Count(line, separators[sepIndex])
		}
	}
	return counts
}

func guessSep(counts [][]int, separators []string, linesRead int) string {
	for sepIndex := range separators {
		same := true
		count := counts[sepIndex][0]
		for lineIndex := 1; lineIndex < linesRead; lineIndex++ {
			if counts[sepIndex][lineIndex] != count {
				same = false
				break
			}
		}
		if count > 0 && same {
			return separators[sepIndex]
		}
	}
	return ""
}

func report(separator string)  {
	switch separator {
	case "":
		fmt.Println("whitespace-separated or not separated at all")
	case "\t":
		fmt.Println("tab-separated")
	default:
		fmt.Printf("%s-separated\n", separator)
	}
}

/*
判断分隔符
/usr/lib/golang/bin/go run /root/github/go/src/guess_separator/main.go /root/github/go/src/guess_separator/in.txt
*-separated
*/
