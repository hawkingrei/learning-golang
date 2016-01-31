// map1
package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "tom", "room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "room 101,..."}
	person, ok := personDB["12345"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	myMap := make(map[string]PersonInfo)
	//myMap = make(map[string] PersonInfo,100)
	myMap = map[string]PersonInfo{"1234": PersonInfo{"1", "jask", "room01,..."}}
	myMap["1234"] = PersonInfo{"1", "jask", "room01,..."}
	fmt.Println(myMap)
	delete(myMap, "1234")
	value, ok := myMap["1234"]
	if ok {
		fmt.Println(value)
	}
}
