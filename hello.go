package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

	// Make a slice first is the TYPE in below case INT
	// The second parameter is the lenght
	// The third parameter here is the cap, ( Maximum number of elements it can hold this can increase if it is exeeded ) basically the
	// Underlaying array is size or allocation is increased.
	tumeric := make([]int, 5, 100)
	fmt.Println(tumeric)
	fmt.Printf("The length of the slice is: %v\n", len(tumeric))
	fmt.Printf("The total amount of elements the Slice can hold: %v\n", cap(tumeric))

	// Composite literal and another slice init example without make --> This is less efficient.
	sli := []int{1, 2, 3, 4, 5}
	fmt.Println(sli)

	// Slicing slices.
	fmt.Println(sli[1:])
	fmt.Println(sli[:4])
	fmt.Println(sli[2:4])

	// Append a slice to another slice.
	slix := []int{50, 60, 70, 80}
	sli = append(sli, slix...)
	fmt.Println(sli)

	// Deleting from a slice, have to slice it a the positions you want and then take out all the data with the ... if not it will give TYPE error.
	sli = append(sli[:3], sli[5:]...)
	fmt.Println(sli)

	// Another way to iterate through a slice ( without using range ) - range is recommended.
	for i := 0; i <= 4; i++ {
		fmt.Println(i, sli[i])
	}

	// Two dimensional slice
	twoD := make([][]int, 3)
	fmt.Println(twoD)

	// Example of Maps which are like dictionaries in python
	m := make(map[string]string)
	m["ED SHERAN"] = "Castle on the hill"
	m["JHON LENON"] = "Imagine"
	fmt.Println(m)

	// Check wheater a key exists in a map
	value, ok := m["Kevin"]
	fmt.Println(value) // Prints out the value of the key kevin if it exists.
	fmt.Println(ok)    // This ( ok ) is a bool that will evaluate to false if it does not exists or true if it does.

	// if check to see if value exist syntax // if kevin is in map then the ok will be true and the if will execute and vise versa.
	if value, ok := m["Kevin"]; ok {
		fmt.Println("This is checking if a map key exists", value)
	}

	// Add a new value and key to the map
	m["Kevin"] = "Portlan main"

	// Iterate through the map and print the keys and values.
	fmt.Print("Printing the keys and values of the map... \n")
	for key, values := range m {
		fmt.Println(key, values)
	}

	// Delete a key from a map, Syntax --> delete(<MapName>, KEY)
	delete(m, "Kevin")
	fmt.Println("This is the new map", m)

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

	// rectangle := rect{width: 10, height: 5}
	// fmt.Println("Area: ", rectangle.area())
	// fmt.Println("Perim: ", rectangle.perim())

	// measure(rectangle)

	// recta := rect{width: 10, height: 5}
	// measure(recta)

	recta := rect{width: 20, height: 40}
	measure(recta)

	area, perim := measure(recta)
	fmt.Printf("The area of the rectangle is: %v and the Perimeter is: %v\n", area, perim)

	f("Direct")
	go f("Goroutine")

	// Example go func
	go func(msg string) {
		fmt.Println(msg)
	}("Starting Parrallel Go routine")

	time.Sleep(time.Second)
	fmt.Println("done")

	// Example of channles. // Create a new channel with make(chan val-type) // Messages can be passed from one go routine to another
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	// TCP port scanner slow
	// for i := 1; i <= 1024; i++ {
	// 	address := fmt.Sprintf("scanme.nmap.org:%d", i)
	// 	conn, err := net.Dial("tcp", address)
	// 	if err != nil {
	// 		// Port is closed or filtered
	// 		continue
	// 	}
	// 	conn.Close()
	// 	fmt.Printf("%d open\n", i)
	// }

	// Example port runner with the waitgroup and sync faster using a worker pool // The worker function is defined outside of main
	// ports := make(chan int, 100)
	// var wg sync.WaitGroup
	// for i := 0; i < cap(ports); i++ {
	// 	go worker(ports, &wg)
	// }
	// for i := 1; i < 1024; i++ {
	// 	wg.Add(1)
	// 	ports <- i
	// }
	// wg.Wait()
	// close(ports)

	// Example fast TCP port scanner and passing results to a result channel and gathering the openports in the openports array then sorting it in order.
	// ports := make(chan int, 100)
	// results := make(chan int)
	// var openports []int
	// for i := 0; i < cap(ports); i++ {
	// 	go worker(ports, results)
	// }

	// go func() {
	// 	for i := 1; i <= 1024; i++ {
	// 		ports <- i
	// 	}
	// }()

	// for i := 0; i < 1024; i++ {
	// 	port := <-results
	// 	if port != 0 {
	// 		openports = append(openports, port)
	// 	}
	// }
	// close(ports)
	// close(results)
	// sort.Ints(openports)
	// for _, port := range openports {
	// 	fmt.Printf("%d open\n", port)
	// }
	// // End of TCP Scanner pool \\

	// // Worker pool example \\
	// const numJobs = 5
	// jobs := make(chan int, numJobs)
	// result := make(chan int, numJobs)

	// // Define how many workers you want
	// for w := 1; w <= 3; w++ {
	// 	go workers(w, jobs, result)
	// }

	// // Pass the job number to numJobs
	// for j := 1; j <= numJobs; j++ {
	// 	jobs <- j
	// }
	// close(jobs)

	// for a := 1; a <= numJobs; a++ {
	// 	<-result
	// }
	fmt.Print("\n")
	sampleBuffchannel()
	// close(messages)
	g, h := split(20)
	fmt.Println(g, h)

	sq := Sqrt(261.1)
	fmt.Printf("The square root of the number is %v and the TYPE is %T\n", sq, sq)

	osv := getOS()
	// Switch with no condition is the same as switch true
	switch osv {
	case "linux":
		fmt.Println("Windows sucks!!!")
	case "darwin":
		fmt.Println("YEY its Mac")
	}

	// A pointer holds the memory address of a value.
	// The & operator generates a pointer to its operand.
	// The * operator denotes the pointer underlaying value.
	pointer, point := 40, 50
	punt := &pointer
	fmt.Println(punt, *punt)
	pont := &point
	fmt.Println(pont, *pont)
	// Change the value of the pointer using *
	*pont = 35
	fmt.Println(point)
	fmt.Println(*pont)

	// file, err := os.Open("README.md")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	buf := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">")
	// 	sentence, err := buf.ReadBytes('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(string(sentence))
	// 	}
	// 	//fmt.Println(scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Println(err)
	// }

	sa1 := secreatagent{
		person1: person1{
			"Kevin",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secreatagent{
		person1: person1{
			"Kelvin",
			"Bond",
		},
		ltk: false,
	}

	sa1.speak()
	sa2.speak()

	// Anonymous function example
	func(x int) { // takes an int
		fmt.Println(x)
	}(42) // Pass in the args

	// Function expressions
	fu := func() {
		fmt.Println("The func expression")
	}
	fu()

	fs := func(x int) {
		fmt.Println("Do you know what year it is: ", x)
	}
	fs(1984)

	vuv := bar()
	fmt.Printf("THE Returned value is of type: %T\n ", vuv)

	yi := incrementor()
	ji := incrementor()

	fmt.Println(yi())
	fmt.Println(yi())
	fmt.Println(yi())
	fmt.Println(ji())
	fmt.Println(ji())

	trk := truck{
		delivers: "Fruits",
	}

	fmt.Println(trk)
	changeMe(&trk)
	fmt.Println(trk)

	// Checking the number of CPU's a machine has
	fmt.Printf("The number of CPU's is: %v\n", runtime.NumCPU())

	// Checking the number of Goroutines
	fmt.Printf("The number of Go Routines is : %v", runtime.NumGoroutine())

	// ================================= END OF MAIN ================================= //
}

// Go only has one looping construct the FOR loop.
// The basic for loop has three components separated by semicolon
// The init statement : Executed before every iteration
// The condition expression : evaluated before every iteration.
// The post statement : Excuted at the end of every iteration.
// The init and post statements are optional.
// The condition is nessesary if is not present it will loop forever.

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("The type OF z is: %T\n", z)
		// fmt.Printf("The type OF x is: %T\n", x)
		// fmt.Println("The value of x is:", x)
		z := z
		// fmt.Println("The value of z is: ", z)
		if z*z == float64(x) {
			// fmt.Println(z * z)
			break
		}
	}
	return z
}

// Function declaration syntax
// func (r receiver) identifier(parameters) (return(s)) { code }

// Anonymous functions --> E.G func(){...}() --> Have to include the param at the end.
// func(x int) {
// 	fmt.Println(x)
// }(42)

// Example Function declaration outside of main
func plus(a int, b int) int {
	return a + b
}

// Example of function with multiple return values
func vals() (int, int) {
	return 3, 7
}

// Example of functions returning functions
func bar() func() int {
	return func() int {
		return 451
	}
}

// Example of function two NAMED values it return.
// Go's returned values may be named. if so they are treated as variables defined at the top of the function.
// This is known as a naked return
// Naked returns should only be used in short functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 2
	y = x - sum
	return
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

type person1 struct {
	first string
	last  string
}

// Create a new function with access to the person struct // In this case its returning the person struct not very usable.
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

// Example of interfaces Go mechanism for grouping and naming related sets of methods. Also any type that has
// the methods in the interface is also of the interface name type.
type geometry interface {
	area() float64
	perim() float64
}

// Go supports methods defined on Structs   // Syntax is func (*struct) funName() type { }
type rect struct {
	width, height float64
}

func (r rect) area() float64 { // Has to be without the (r *rect) to be able to implement the methods in an interface.
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Measure implements the methods stored in geometry interface.
func measure(g geometry) (float64, float64) {
	// fmt.Println(g)
	// fmt.Println(g.area())
	// fmt.Println(g.perim())
	return g.area(), g.perim()
}

// Example of GO Routines  // A GO routine is a lightweight thread of execution. // Go Routines can be called by the go keyword example funciton in main go func
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// func worker(ports chan int, wg *sync.WaitGroup) {
// 	for p := range ports {
// 		fmt.Println(p)
// 		wg.Done()
// 	}
// }

// Example of worker function using another channel for the results.
// func worker(ports, results chan int) {
// 	for p := range ports {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", p)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			results <- 0
// 			continue
// 		}
// 		conn.Close()
// 		results <- p
// 	}
// }

// // Example of Another worker function
// func workers(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("worker", id, "started job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finish job", j)
// 		results <- j * 2
// 	}
// }

// Simple example of buffered channels
// By default channels are unbuffered meaning that they will only accept sends ( chan <- )
// if there is a corresponding receive ( <- chan ) ready to received the sent value
//Buffered channel accept a limited number of values without a corresponing receiver for those values.

// Example channel declaration send/receive/bidirectional
// c := make( <- chan int)  // this is a receive only channel
// cs := make(chan<- int) // Send only channel
// cs := make(chan int) // Bidirectional channel
// bc := make(chan int, 5) // Example Bidirectional Buffered channel can send and receive five things at once.

func sampleBuffchannel() {
	messages := make(chan string, 5) // Five is the maximum things the channel can hold at one.
	// If for exmple more than three strings are passed at once it will crash.
	// For loop passes on a the time in this case.
	// messages <- "Buffered"
	// messages <- "channel"
	// messages <- "Third"
	words := make([]string, 5)
	words[0] = "Buffered"
	words[1] = "channel"
	words[2] = "ThirdandFive"
	words[3] = "ThirdandFour"
	words[4] = "ThirdandSix"
	// fmt.Println(words)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	for _, word := range words {
		messages <- word
		// time.Sleep(time.Second)
		// fmt.Printf("Hello message: %v \n", <-messages)
	}
	close(messages)
	fmt.Println("########################## Printing messages from the channel ##########################")
	fmt.Print("\n")
	for elem := range messages {
		time.Sleep(time.Second)
		fmt.Printf("Hello message: %v \n", elem)
		fmt.Print("\n")
		// close(messages)
	}
	fmt.Println("################### Done ###################")

}

// getOS() provides the current user os enviroment
func getOS() string {
	// Get the current os
	osv := runtime.GOOS
	return osv
}

// Methods for types.
type secreatagent struct {
	person1 // Here this is of the person1 struct type .
	ltk     bool
}

// We can now declare a function with a receiver of the secret  agent type and all types of secret  agent will have access to this method.
func (s secreatagent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// Closure functions --> Below example returns a func that returns an int
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

type truck struct {
	delivers string
}

// The changeMe functions take a pointer to truck, can be called by changeMe(&truck)
func changeMe(p *truck) {
	p.delivers = "I can deliver everything and anything"
}

// Json marshall example --> https://play.golang.org/p/HINWtn30TsG
// Json Unmarshal example --> https://play.golang.org/p/UPCdkPKLPmC
// Concurrency example --> https://play.golang.org/p/KbwIAtNuy2c
