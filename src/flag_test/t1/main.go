package main

import (
	"flag"
	"fmt"
	"os"
)

type TestOpts struct {
	opt1 string
	opt2 string
	opt3 bool
}

var (
	opts TestOpts
	testFlags = flag.NewFlagSet("mhc", flag.ExitOnError)
)

func init() {
	testFlags.StringVar(&opts.opt1, "name1", "value1", "1111")
	testFlags.StringVar(&opts.opt2, "name2", "value2", "2222")
	testFlags.BoolVar(&opts.opt3, "name3", false, "333")

	testFlags.Parse(os.Args[1:])
}

func main() {
	//a
	testFlags.Visit(func(f *flag.Flag) {
		fmt.Printf("visit flag %s\n", f.Name)
	})
	testFlags.VisitAll(func(f *flag.Flag) {
		fmt.Printf("visit all flag: %s\n", f.Name)
		err := testFlags.Set(f.Name, "222")
		if err != nil {
			fmt.Printf("flag %s err: %s\n", f.Name, err)
		}
	})
	fmt.Println(opts.opt1)
	fmt.Println(opts.opt2)
	fmt.Println(opts.opt3)
}
