package main

import (
	"container/list"
)

func main(){
	type entry struct {
		key   interface{}
		value interface{}
		mp    map[string]interface{}
	}
	
	
	var mylist *list.List
	var mymap map[string]*entry
	mymap = make(map[string]*entry)	
	mylist = list.New()
	ent := &entry{"wwz","abc",make(map[string]interface{})}
	mymap["wwz"] = &entry{"wwz","abc",make(map[string]interface{})}
	(*mymap["wwz"]).mp["wwz"] = list.New()
	mylist.PushFront(&ent)
	
	
}
