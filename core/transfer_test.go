package core

import (
	"fmt"
	"testing"
	"time"
)

func TestTransfer(t *testing.T){
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
	want1 := `{"$.key1":1,"$.key2":2,"$.key3.key4":4,"$.key3.key5":5,"$.key6[0].key7":7,"$.key6[1].key8":8,"$.key6[2][0].key9":9}`
	result,err := TransferToOneLevel(testJson)
	if err != nil && result != want1 {
		t.Errorf("TransferToOneLevel fail res=%s,want=%s",result,want1)
	}

	want2 := `{"$":{"key1":1,"key2":2,"key3":{"key4":4,"key5":5},"key6":[{"key7":7},{"key8":8},[{"key9":9}]]},"$.key1":1,"$.key2":2,"$.key3":{"key4":4,"key5":5},"$.key3.key4":4,"$.key3.key5":5,"$.key6":[{"key7":7},{"key8":8},[{"key9":9}]],"$.key6[0]":{"key7":7},"$.key6[0].key7":7,"$.key6[1]":{"key8":8},"$.key6[1].key8":8,"$.key6[2]":[{"key9":9}],"$.key6[2][0]":{"key9":9},"$.key6[2][0].key9":9}`
	result2,err2 := TransferToOneLevelShowAll(testJson)
	if err2 != nil && result2 != want2 {
		t.Errorf("TransferToOneLevelShowAll fail res=%s,want=%s",result2,want2)
	}
}

func TestBanch(t *testing.T){
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
	startTime := time.Now()
	for i:=0;i<100000;i++ {
		TransferToOneLevel(testJson)
	}
	fmt.Println("TransferToOneLevel 100000 times cost",time.Since(startTime))

	startTime2 := time.Now()
	for i:=0;i<100000;i++ {
		TransferToOneLevelShowAll(testJson)
	}
	fmt.Println("TransferToOneLevel 100000 times cost",time.Since(startTime2))

}
