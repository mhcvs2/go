package main

import (
	"text/tabwriter"
	"os"
	"fmt"
)

func t11() {
	w := tabwriter.NewWriter(os.Stdout, 0, 7, 11, '*', 0)
	defer w.Flush()
	fmt.Fprint(os.Stdout, "aa\tbb\n")
	fmt.Println("------------------")
	fmt.Fprint(w, "aa\tbb\n")
}
//aa	bb
//------------------
//aa***********bb

func main() {
	t11()
}
