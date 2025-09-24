package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

type Todo struct {
    ID         string `json:"id"`
    Todo       string `json:"todo"`
    IsComplete bool   `json:"is_complete"`
}


func main(){

	http.HandleFunc("/",serverHealth)
	http.HandleFunc("/create-todo",createTodo)
	http.HandleFunc("/get-todos",getTodo)
	http.HandleFunc("/update-todos",updateTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	
}

func serverHealth(w http.ResponseWriter,r *http.Request){

	resposne := Response {
		Error: false,
		Message: "Server is up!",
	}

	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(resposne)
}

func createTodo(w http.ResponseWriter,r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")	
		json.NewEncoder(w).Encode(Response {Error: true,Message: "Method not allowed!",})
		return
	}

	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.Header().Set("Content-Type", "application/json")	
		json.NewEncoder(w).Encode(Response {Error: true,Message: "Invalid JSON!"})
		return 
	}

	saveTodo(todo.Todo)

	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(Response {Error: false,Message: "Todo created successfully!",})
}

func getTodo(w http.ResponseWriter,r *http.Request){

	var allTodos []Todo = getAlltodos()

	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(allTodos)
}

func updateTodo(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	updateTodoStatus(id)
	json.NewEncoder(w).Encode(Response {Error: false,Message: "Todo updated successfully!",})

}