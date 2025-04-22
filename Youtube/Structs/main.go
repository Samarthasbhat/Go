package main

// struct is comparable if all its fields are comparable
// structs are "passed by value" unless a pointer is used

import (
	"encoding/json"
	"fmt"
	"time"
)
type Response struct{
	Page int `json:"page"`
	Words []string `json:"words,omiempty"`
}


type Employee struct{
	Name string
	Number int
	Boss *Employee
	Hired time.Time
}

type album1 struct{
	title string
}

type album2 struct{
	title string
}


func main() {

	c := map[string]*Employee{}  

	c["Lamine"] = &Employee {"Lamine", 2, nil, time.Now()}
	c["Lamine"].Number++

	 c["Matt"] = &Employee {
		Name: "Matt",
		Number: 1,
		Boss: c["Lamine"],
		Hired: time.Now(),
	}

	// e.Name = "Matt"
	// e.Number = 1
	// e.Hired = time.Now()



	fmt.Printf("%T %+[1]v\n",  c["Lamine"])
	fmt.Printf("%T %+[1]v\n",  c["Matt"])


	var a1 = album1{
		"The white Album",
	}
	

	var a2 = album2{
		"The Black Album",
	}

	a1 = album1(a2)
	

	fmt.Println(a1, a2)


	v1  := struct {
		X int `json: "foo"`
	}{1}

	v2  := struct {
		X int `json: "foo"`
	}{2}

	v1 = v2
	fmt.Println(v1)

	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n",r)

	var r2 Response

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n",r)

}

