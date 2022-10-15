package group

import (
	"fmt"
)

func PrettyPrint(sourceStr str, targetStr str) {
	fmt.Printf("\"%s\" > \"%s\"\n", sourceStr, targetStr)
}