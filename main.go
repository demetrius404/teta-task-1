package main

import (
	"fmt"
	lib "local/pkg"
)

func main() {

	fmt.Println("--- Dive into Golang 2.0 ---")
	sourceStr := "zzzzcccUUUuu"
	lib.PrettyPrint(sourceStr, lib.GroupStr(sourceStr))
	sourceStr = "aaabbbccccc"
	lib.PrettyPrint(sourceStr, lib.GroupStr(sourceStr))
	sourceStr = "cccccaaabbb"
	lib.PrettyPrint(sourceStr, lib.GroupStr(sourceStr))
}