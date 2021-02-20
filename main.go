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
	},
	"key6": [{
		"key7": 7
	}, {
		"key8": 8
	},[
		{
			"key9":9	
		}
	]]
}`
	result,err := core.TransferToOneLevel(testJson)
	fmt.Println(err)
	fmt.Println(result)


	result2,err2 := core.TransferToOneLevelShowAll(testJson)
	fmt.Println(result2)
	fmt.Println(err2)
}
