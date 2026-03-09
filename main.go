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
}
