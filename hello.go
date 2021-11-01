package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("The value of i now is %v\n", i)
	}

	x := "Pizza"
	var y string
	var z int
	const n int = 5
	if x == "Yummy" {
		fmt.Printf("The food is %v\n", x)
	} else {
		fmt.Println("The food is not Yummy")
	}

	if x != " " {
		fmt.Printf("x does exist and it's not empty!\n")
	}

	if y != "" {
		fmt.Printf("Y is not an empty string %v", y)
	} else {
		fmt.Println("Y is an empty string")
	}

	if z == 0 {
		fmt.Printf("Z is an int with the value %v!!!\n", z)
	} else {
		fmt.Printf("Z is not an int with the value of zero\n")
	}

	if n == 5 {
		fmt.Printf("N is a constant int with a value of %v which does not change!!\n", n)
	} else {
		fmt.Println("N is some weird value")
	}
	switch x {
	case "Yum":
		fmt.Printf("This is the word %v\n", x)
	case "Yummy":
		fmt.Printf("This is the word %v\n", x)
	case "Pizza":
		fmt.Printf("This is the word:  %v\n", x)
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	osPath, error := os.Getwd()
	if error != nil {
		fmt.Println(error)

	} else {
		fmt.Printf("This is the current Path: -- %v\n", osPath)
	}

	hostname, error := os.Hostname()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("The box hostname is: ######### %v #########\n", hostname)
	}

	userHomeDirectory, error := os.UserHomeDir()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("The User home is: ######### %v #########\n", userHomeDirectory)
	}
	filePath := filepath.Join(osPath, "New_dir")
	fmt.Println(filePath)

	err := os.Mkdir(filePath, 0777)
	if err != nil {
		fmt.Println(err)
	}

	prob := os.WriteFile("New_dir/text.txt", []byte("Hello, Gophers!"), 0666)
	if prob != nil {
		fmt.Println(prob)
	}

	files, error := os.ReadDir(".")
	if error != nil {
		fmt.Println(error)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println(file.Name())
		}
	}
	current_time := time.Now()
	formatted_time := current_time.Format("2006-01-02")
	fmt.Printf("Today's date is: %v\n", formatted_time)

	//os.Chdir(filePath) Changes to the directory given if possible
	//os.Mkdir(formatted_time, 0777)
	//fmt.Println(os.Environ()) Way to get all enviroment variables.

	// Array example
	var arr [5]int          // Example declaration
	for _, s := range arr { // Range can be used to iterate over arrays as well as maps
		fmt.Printf("%v\t", s) // Got to print and set both the the position and the value or throw away the position of the array.
	}
	arr[0] = 100
	arr[1] = 500
	fmt.Printf("\n%v\t%v\n", arr[0], arr[1])

	// Slice examples
	slice := make([]string, 5)
	fmt.Println(slice)
	// Append to the slice and add, if you append it will not replace the five empty possitions at the start.
	slice[0] = "World Peace"
	slice[1] = "Knocking on Heavens door"
	slice[2] = "The Parting Glass"
	slice[3] = "Photograph"
	slice[4] = "Castle on the hill"
	slice = append(slice, "Robert Moss", "Ed Sheran", "Black Eyed Pays", "Legend", "Lenon", "RBD", "Backstreet Boys")
	fmt.Println(slice)

	twoD := make([][]int, 3)
	fmt.Println(twoD)

	// Example of Maps which are like dictionaries in python
	m := make(map[string]string)
	m["ED SHERAN"] = "Castle on the hill"
	m["JHON LENON"] = "Imagine"
	fmt.Println(m)
	// Declare a map in a single line
	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", newMap)

	sum := plus(750, 750)
	fmt.Println(sum)

	r, v := vals()
	fmt.Println(v)
	fmt.Println(r)

	_, o := vals()
	fmt.Println(o)

	tottal := []int{1, 2, 3, 4, 43}
	addition(tottal...)

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{})
	Kevin := newPerson("Kevin")
	Kevin.age = 50
	fmt.Println(Kevin)

	Ali := student("Ali", 11, 6, 7, 4.0, "Harvard", "Fotball")
	fmt.Println(Ali.name)
	fmt.Println(Ali.current_grade)
	fmt.Println(Ali.gpa)
	fmt.Println(Ali.club)
	fmt.Println(Ali.school)
	Ali.school = "Yale"
	fmt.Println(Ali.school)
	fmt.Println(&Ali.school) // This references the value set in the Student struct which is a space set in memory for a string

	// ================================= END OF MAIN ================================= //
}

// Example Function declaration outside of main
func plus(a int, b int) int {
	return a + b
}

// Example of function with multiple return values
func vals() (int, int) {
	return 3, 7
}

// Variadic Functions can be used with any numbers of trailing arguments.
// The function below will take any number of ints as arguments.

func addition(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Example of Structs
type person struct {
	name string
	age  int
}

// Create a new function with access to the person struct
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

//Student struct
type Student struct {
	gpa           float64
	name          string
	age           int
	school        string
	club          string
	current_grade [2]int
}

//Create a new student and implement Student
func student(name string, age int, current_grade int, grade int, gpa float64, school string, club string) *Student {
	stu := Student{name: name}
	stu.age = age
	stu.current_grade[0] = current_grade
	stu.current_grade[1] = grade
	stu.gpa = gpa
	stu.school = school
	stu.club = club
	return &stu

}
