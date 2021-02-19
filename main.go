package main

import (
	"fmt"
	"github.com/transferOneLevelJson/core"
)

func main(){
	testJson := `{
	"key1": 1,
	"key2": 2,
	"key3": {
		"key4": 4,
		"key5": 5
	}
}`
	result,err := core.TransferToOneLevel(testJson)
	fmt.Println(err)
	fmt.Println(result)
}