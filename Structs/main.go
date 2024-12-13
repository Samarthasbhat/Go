package main

import ("fmt")

type Person struct{
	Name string
	Age int
	Gender string
	Phone string
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

}


