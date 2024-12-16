package main

import (
	"fmt"
)


type Student struct{
	ID int
	Name string
	Courses []string //List of Courses
	Grades map[string]string // Map to store grades for each course
}

//Predefined array of courses
var courses = [3]string{"Mathematics", "Physics", "Computer Science"}

func main(){
	// Slice to store all students

	var students []Student

	//Main menu 
	for{
		fmt.Println("\nStudent Management System")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Add Grades")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice{
		case 1:
			addStudent(&students)
		case 2:
			viewStudents(students)
		case 3:
			addGrades(&students)
		case 4:
			fmt.Println("Exiting.....")
			return
		default:
			fmt.Println("Invalid choice! Please try again..")
		}
	}
}


//Function to add a new student
func addStudent(students *[]Student){
	var name string
	var id int

	fmt.Print("Enter Student ID:")
	fmt.Scan(&id)
	fmt.Print("Enter the student name:")
	fmt.Scan(&name)
	
	//  Create a new student and add predefined courses

	newStudent := Student{
		ID: id,
		Name: name,
		Courses: courses[:], // Slicing the array to initialize the courses slice
		Grades: make(map[string]string),
	}

	*students = append(*students, newStudent)
	fmt.Println("Student added successfuly!")
}

//Function to view all students
func viewStudents(students []Student){
	if len(students) == 0{
		fmt.Println("No students found!")
		return
	}

	fmt.Println("\n List of students:")

	for _, student := range students {
		fmt.Printf("ID: %d, Name:%s, Courses:%v\n",student.ID, student.Name, student.Courses)


		if len(student.Grades) > 0{
			fmt.Println("Grades:")
			for course, grade := range student.Grades{
				fmt.Printf(" %s: %s\n", course, grade)
			}
		}else{
			fmt.Println("No grades added yet.")
		}
	}
}

// Function to add grades for a student
func addGrades(students *[]Student){
	if len(*students) == 0{
		fmt.Println("No students available to add grades!")
		return
	}
	

	var id int
	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)
	
	// Find the student by ID
	var student *Student
	for i := range *students{
		if(*students)[i].ID == id{
			student = &(*students)[i]
			break
		}
	}
	// if student != student.id { 
	// 	fmt.Println("There is no student present")
	// }

	if student == nil{
		fmt.Println("Student not found!")
		return
	}

	// Add grades for each course

	for _,course := range student.Courses{
		var grade string
		fmt.Printf("Enter grade for %s: ", course)
		fmt.Scan(&grade)
		student.Grades[course] = grade
	}

	fmt.Println("Grades added successfully!!")
}