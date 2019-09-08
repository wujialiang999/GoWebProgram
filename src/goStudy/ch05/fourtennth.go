package main

import "sync"

//延迟并发解锁
func main(){

}
var (
	valueByKey =make(map[string] int)
	valueByKeyGuard sync.Mutex
)

//根据键读取值
func readValue(key string) int{
	valueByKeyGuard.Lock()
	defer  valueByKeyGuard.Unlock()
	return valueByKey[key]
}
