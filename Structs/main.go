package main

import ("fmt")

type Person struct{
	Name string
	Age int
	Gender string
	Phone string
}

// Structs with Nested Fields

type Address struct {
	City string
	Zipcode  string
}

type Human struct{
	Name string
	Address Address
}

func main(){

person1 := Person{
	Name: "John",
	Age: 40,
	Gender: "Male",
	Phone: "9237864562",
}

// without specifying all fields
person2 := Person{Name: "Alice", Gender: "Female"}



// Access fields

fmt.Println("Person 1 :", person1)
fmt.Printf("Name: %s, Phone: %s\n", person1.Name, person1.Phone)

    // Modify fields
    person2.Phone = "6363271234"
    fmt.Printf("Person 2: %+v\n", person2)

	//display the person
	fmt.Println("Before update:")
	displayPerson(person1)
	
	//Update the phone number
	updatePhone(&person1, "8237652331")

	fmt.Println("After Update:")
	// Display the updated person
	displayPerson(person1)


	// Usage for nested
	person := Human{
		Name: "Ron",
		Address: Address{
			City: "New York",
			Zipcode: "10001",
		},
	}
	fmt.Printf(" Nested struct \n City: %s, Zipcode: %s\n", person.Address.City, person.Address.Zipcode)
}

// Pass by value

func displayPerson(p Person){
	fmt.Printf("Name: %s, Age: %d, Phone:%s\n", p.Name, p.Age, p.Phone)
}

func updatePhone(p *Person, newPhone string){
	p.Phone = newPhone // Modify the field via pointer
}


