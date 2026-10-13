package main

import "fmt"

type person struct {
	name      string
	birthDate string
}

type student struct {
	person
	course string
}

func main() {
	e1 := student{
		person: person{
			name:      "João",
			birthDate: "2000-01-01",
		},
		course: "Engenharia",
	}

	e2 := student{person{"Maria", "2002-05-15"}, "Medicina"}
	fmt.Println(e2)

	fmt.Println("Nome:", e1.name)
	fmt.Println("Nascimento:", e1.birthDate)
	fmt.Println("Curso:", e1.course)
}
