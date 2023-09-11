package main

import (
	"fmt"
)

type UserProfile struct {
	name    string
	age     int
	city    string
	country string
}

// displayAddress() method uses User as the receiver type
func (user UserProfile) displayAddress() {
	fmt.Printf("Address of %s is %s, %s \n", user.name, user.city, user.country)
}

func main() {
	user := UserProfile{
		name:    "Amr",
		age:     30,
		city:    "Cairo",
		country: "Egypt",
	}
	user.displayAddress() //Call displayAddress() method of User type
}
