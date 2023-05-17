package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// first way to make a struct variable
	alex := person{"Alex", "Anderson", contactInfo{"alex@mail.com", 48325}}
	// second way to make a struct variable
	jane := person{
		firstName: "Jane",
		lastName:  "Henderson",
		contact: contactInfo{
			email:   "jane@mail.com",
			zipCode: 92849,
		}}
	// third way to make a struct variable
	var igdi person
	igdi.firstName = "Igdi"
	igdi.lastName = "Barnes"
	var igdiContact contactInfo
	igdiContact.email = "igdi@mail.com"
	igdiContact.zipCode = 93048
	igdi.contact = igdiContact

	igdi.updateName("Mr. Wiggles")
	igdi.print()
	alex.print()
	jane.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
