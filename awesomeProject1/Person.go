package main

import "fmt"

type Person struct {
	name string
}

func (person Person) handleEvent(vacancies []string) {
	//TODO implement me
	fmt.Println("Hi, dear", person.name)
	fmt.Println("Vacancies updated:")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}
