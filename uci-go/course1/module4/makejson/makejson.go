package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Print("Enter a name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	var address string
	fmt.Print("Enter an address: ")
	fmt.Scanln(&address)
	address = strings.TrimSpace(address)
	
	var user = map[string]string {"name": name, "address": address}
	buser, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", buser)
}