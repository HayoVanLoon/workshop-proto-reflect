package main

import (
	"fmt"
	"os"

	"workshop/ex1map"
	"workshop/ex2tree"
	"workshop/ex3reflect"
	"workshop/ex4proto"
)

func printHelp() {
	fmt.Printf("Usage:\n\t%s map|tree|reflect|proto", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(3)
	}
	switch os.Args[1] {
	case "map":
		fmt.Println("==== with map ====")
		ex1map.Run()
	case "tree":
		fmt.Println("==== with tree ====")
		ex2tree.Run()
	case "reflect":
		fmt.Println("==== with reflect ====")
		ex3reflect.Run()
	case "proto":
		fmt.Println("==== with proto reflect ====")
		ex4proto.Run()
	default:
		printHelp()
		os.Exit(3)
	}
	fmt.Println()
}
