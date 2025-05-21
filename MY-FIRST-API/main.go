package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
you can get directly this structly just copying json from diffrent source

	to here. It will ask you to generate Go Type.
	In GO "struct" are stuctures are user-defined types, that we used to store
	different collection of fields into a single field
	you can thinks its like a data structurehat group together,
*/
type TodoItem struct {
	Item string `json:"item"`
}

// survmux is a server
// we're going to attach HTTP endpoints to this mux, then we will be able to receive our requests
func main() {
	//its a slice not an array
	var todos = make([]string, 0)
	mux := http.NewServeMux()

	//endpoint
	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		//below line taking the response Write from function,
		//and its writing Hello World back to the user
		_, err := w.Write([]byte("Hello World"))
		//below line error handling
		if err != nil {
			// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
			// it make our application stop
			log.Fatal(err)
			//but it not going to stop because this is how mux works
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)
		//if request doesnt make this type it send bad request status
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		//here you are appending the todo item and assigning back to the todos
		todos = append(todos, t.Item)
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /all-todos", func(w http.ResponseWriter, r *http.Request) {
		//json.Marshal(v any) takes anything, we use this line of code to convert our todos slice to byte[]
		//because Write function only write byte[]
		b, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
		}
		//you are seeing this _ again and again its means you actuallu have a value, but its says i am not interested in that, you can _ underscor for error also
		//below line writing back to user todods as byte because Write function write only Byte[]
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})
	//if you're getting autocomplete its because Golang has local LLM to help you complete the code.

	//setting up the mux to port 8080
	//if for any reason our program stop it return error, you can see in the console
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
