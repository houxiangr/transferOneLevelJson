package core

import "sync"

//todo 等go-cache完成后替换
var cacheMap map[string]string
var cacheLock sync.RWMutex

func init(){
	cacheMap = make(map[string]string)
	cacheLock = sync.RWMutex{}
}
