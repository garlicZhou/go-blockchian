package main

import (
	"flag"
	"fmt"
)

var species = flag.String("species", "go", "the usage of flag")

var num = flag.Int("ins",1, "ins nums")

func main()  {
	flag.Parse()
	fmt.Println("a string flag", *species)
	fmt.Println("ins num:", *num)
}

