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
	printloops()
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
