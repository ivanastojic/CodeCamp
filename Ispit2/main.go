package main

import (
	"fmt"
)

type Employee struct {
	Name     string
	Position string
	Tasks    []string
}

func main() {

	zaposlenik1 := Employee{
		Name:     "Kata Lukić",
		Position: "računovodja",
		Tasks:    []string{"daje place"},
	}

	zaposlenik2 := Employee{
		Name:     "Iva Šarić",
		Position: "menadžer",
		Tasks:    []string{"upravlja svime"},
	}

	zaposlenik1.AddTask("ulazne i izlazne fakture")
	zaposlenik2.AddTask("vodi pr")
	zaposlenik1.AddTask("financije")
	zaposlenik2.AddTask("daje ideje")

	fmt.Println(zaposlenik1.Name)
	fmt.Println(zaposlenik1.Position)
	fmt.Println(zaposlenik1.Tasks)
	fmt.Println(zaposlenik2.Name)
	fmt.Println(zaposlenik2.Position)
	fmt.Println(zaposlenik2.Tasks)

	zaposlenik1.CompleteTasks()
	zaposlenik2.CompleteTasks()
	zaposlenik1.CompleteTasks()
	zaposlenik2.CompleteTasks()

}

func (e *Employee) AddTask(task string) {
	e.Tasks = append(e.Tasks, task)
}

func (e *Employee) CompleteTasks() {
	task := e.Tasks[0]

	fmt.Printf("Employee <%s> completed task: <%s>.\n", e.Name, task)

	e.Tasks = e.Tasks[1:]

}
