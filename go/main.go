package main

import (
	"fmt"
	"os"
	"workshop/withmap"
	"workshop/withproto"
	"workshop/withreflect"
	"workshop/withtree"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "map":
		fmt.Println("==== with map ====")
		withmap.Run()
	case "struct":
		fmt.Println("==== with tree ====")
		withtree.Run()
	case "reflect":
		fmt.Println("==== with reflect ====")
		withreflect.Run()
	case "proto":
		fmt.Println("==== with proto reflect ====")
		withproto.Run()
	case "annotations":
		fmt.Println("==== apply annotations ====")
		apple := withproto.Create()
		withproto.Apply(apple)
		fmt.Println(apple)
	}
	fmt.Println()
}
