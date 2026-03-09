package main

import "fmt"

/* 🧩 Practice 1 — Variables + Basic Logic

Write a Go program that:

Declare:

name (string)

age (int)

Print:

My name is <name> and I am <age> years old

Then check:

If age >= 18
→ print "Adult"

Else
→ print "Minor" */

func main() {
	name := "Amodinee"
	age := 23

	fmt.Println("My name is", name, "and I am", age, "years old")

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// a := 10
	// fmt.Println("Initial value:", a)
	// changenum(&a)
	// fmt.Println("New value:", a)
	//printloops()
	//arrayPractice()
	//slicePractice()
	//mapPractice()
	//prinit()
	// a, b := 5, 10
	// fmt.Println("Before Swap a:", a, "b:", b)
	// swap(&a, &b)
	// fmt.Println("After Swap a:", a, "b:", b)

	// s1 := Student{
	// 	name:  "Amodinee",
	// 	marks: 90,
	// }
	// s2 := Student{
	// 	name:  "Rahul",
	// 	marks: 34,
	// }
	// s1.checkmarks()
	// s2.checkmarks()

	s := newStudent("Amodinee", -90)
	s.Show()
}

/* Write a Go program that:1)Prints numbers 1 to 10 2)Prints only even numbers from 1 to 20
3)Calculates sum of numbers from 1 to 5 */
func printloops() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println() // newline

	fmt.Println("Only Even Numbers:")

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	fmt.Println() // newline

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}

	fmt.Println("Sum is:", sum)
}

func arrayPractice() {
	a := [5]int{10, 20, 30, 40, 50}
	fmt.Println(a)

	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	fmt.Println("Sum is:", sum)
}

func slicePractice() {
	a := []int{10, 20, 30}
	a = append(a, 40)
	a = append(a, 50)

	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("Length:", len(a))
	fmt.Println("Capacity", cap(a))
}

func mapPractice() {
	m := make(map[string]int)
	m["Amodinee"] = 90
	m["Rahul"] = 80
	m["Neha"] = 85

	for k, v := range m {
		fmt.Println(k, v)
	}

	value, ok := m["Rahul"]
	if ok {
		fmt.Println("Rahul found with marks", value)
	} else {
		fmt.Println("Rahul not found")
	}
}

func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}
func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}
func prinit() {
	result1 := operate(3, 5, add)
	result2 := operate(3, 5, multiply)
	fmt.Println("Add:", result1)
	fmt.Println("Multiply:", result2)
}

//1. Write a function that increases a number by 1 using a pointer.
//Call the function and show that the original number changes.

func changenum(ptr *int) {
	(*ptr)++
}

/* Write a function that swaps two numbers using pointers.
Requirements:Create two variables in main
Print their values before swapping
Call a function that swaps them using pointers
Print their values after swapping
Expected behavior example:
Before swap
5 10
After swap
10 5 */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

type Student struct {
	name  string
	marks int
}

func (s Student) checkmarks() {
	if s.marks >= 40 {
		fmt.Println(s.name, "is passed")
	} else {
		fmt.Println(s.name, "is failed")
	}
}

func newStudent(n string, m int) *Student {
	if m < 0 {
		m = 0
	}
	s := Student{
		name:  n,
		marks: m,
	}
	return &s
}

func (s Student) Show() {
	fmt.Println("Name:", s.name)
	fmt.Println("Marks:", s.marks)
}
