package main

import (
	"encoding/json"
	"os"
	"github.com/google/uuid"
)

var filename  string = "todo.json"

func saveTodo(todo string) Todo {
	t := Todo{
		ID:        uuid.New().String(),
		Todo:      todo,
		IsComplete: false,
	}

	t.updateTodoFile()

	return t
}

func (t Todo) updateTodoFile() {

	var todos []Todo
	file, err := os.Open(filename)

	if err == nil {
		defer file.Close()
		//return  err
	}

	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		// ignore empty file or bad JSON
		todos = []Todo{}
	}

	todos = append(todos,t)
	data, _ := json.MarshalIndent(todos,""," ")

	os.WriteFile(filename, data, 0644)
}

func getAlltodos() []Todo{

	var todos []Todo
	//var structurdTodos []Todo

	file, err := os.Open(filename)

	if err == nil {
		defer file.Close()
		//return  err
	}

	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		// ignore empty file or bad JSON
		todos = []Todo{}
	}

	// for _,v := range todos {

	// 	structurdTodos = append(structurdTodos, Todo{
	// 		ID: v.ID,
    //         Todo:       v.Todo,
    //         IsComplete: v.IsComplete,
    //     })

	// }

	return todos
	
}

func updateTodoStatus(id string){

	var todos []Todo

	file, _ := os.Open(filename)

	if err := json.NewDecoder(file).Decode(&todos); err != nil {
		return
	}

	for i,v := range todos {
		if v.ID == id {
			todos[i].IsComplete = true
		}
	}

	data, _ := json.MarshalIndent(todos,""," ")
	os.WriteFile(filename, data, 0644)
}