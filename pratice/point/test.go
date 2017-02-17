package main

import (
	"container/list"
)

func main(){
	type entry struct {
		key   interface{}
		value interface{}
	}
	
	
	var mylist *list.List
	mylist = list.New()
	ent := &entry{"wwz","abc"}
	mylist.PushFront(&ent)
	
	
}
