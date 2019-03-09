package main

import (
	"flag"
	"fmt"
	"os"
	"study/lib"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	flag.StringVar(&name, "name", "everyone", "The greeting object")
}

func main()  {
	flag.Parse()
	lib.Hello(name)
	//fmt.Printf("Hello, %s!\n", name)
}