package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// embed a struct type as a property of another struct
	contact contactInfo
	// do not have to specify field name if we don't want to
	// contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// approach 1
	// tim := person{"Tim", "Thomas"}
	// fmt.Printf("%+v", tim)

	// approach 2
	george := person{firstName: "George", lastName: "Giant"}
	fmt.Printf("%+v \n", george)

	//approach 3∏
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex@gmail.com"
	alex.contact.zipCode = 1234
	fmt.Printf("%+v \n", alex)

    //approach 4
    jim := person{
        firstName: "Jim",
        lastName: "Smith",
        contact: contactInfo{
            email: "jim@gmail.com",
            zipCode: 1234,
        },
    }

   jim.updateName("Jimmy")
   jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
    (*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v \n", p)
}