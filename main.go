package main

import(
"fmt"
"os"
"encoding/json"
)

type Address struct{
	City string
	State string
	Country string
	Pincode json.Number
}

type user struct{
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}


func main(){
	
}