package main

import (
	"fmt"
	"os"

	"workshop/ex1map"
	"workshop/ex2tree"
	"workshop/ex3reflect"
	"workshop/ex4proto"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "map":
		fmt.Println("==== with map ====")
		ex1map.Run()
	case "struct":
		fmt.Println("==== with tree ====")
		ex2tree.Run()
	case "reflect":
		fmt.Println("==== with reflect ====")
		ex3reflect.Run()
		apple := ex3reflect.Create()
		ex3reflect.Apply(&apple)
		fmt.Println(apple)
	case "proto":
		fmt.Println("==== with proto reflect ====")
		ex4proto.Run()
	case "annotations":
		fmt.Println("==== apply annotations ====")
		apple := ex4proto.Create()
		ex4proto.Apply(apple)
		fmt.Println(apple)
	}
	fmt.Println()
}
