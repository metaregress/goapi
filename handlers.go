package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	var todos Todos
	for i, t := range todoRepo {
		//I don't know how to put this in the for body so I'm hacking in here
		if i > currentId-5 { //only give the 5 latest todos
			todos = append(todos, t)
		}

	}
	// todos = append(todos, RepoFindTodo(1))
	// todos = append(todos, RepoFindTodo(2))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(todos)

	if err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoIdString := vars["todoId"]
	todoId, strConvErr := strconv.Atoi(todoIdString)
	if strConvErr != nil {
		fmt.Fprint(w, "Error converting given id to int")
	}
	todo := RepoFindTodo(todoId)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todo)
	if err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "Todo show: %v", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	bodyCloseErr := r.Body.Close()
	if bodyCloseErr != nil {
		panic(bodyCloseErr)
	}
	jsonUnmarshalErr := json.Unmarshal(body, &todo)
	if jsonUnmarshalErr != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		errError := json.NewEncoder(w).Encode(jsonUnmarshalErr)
		if errError != nil {
			panic(errError)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	jsonWriteErr := json.NewEncoder(w).Encode(t)
	if jsonWriteErr != nil {
		panic(jsonWriteErr)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoIdString := vars["todoId"]
	todoId, strConvErr := strconv.Atoi(todoIdString)
	if strConvErr != nil {
		fmt.Fprint(w, "Error converting given id to int")
	}
	RepoDestroyTodo(todoId)
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	bodyCloseErr := r.Body.Close()
	if bodyCloseErr != nil {
		panic(bodyCloseErr)
	}
	jsonUnmarshalErr := json.Unmarshal(body, &todo)
	if jsonUnmarshalErr != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		errError := json.NewEncoder(w).Encode(jsonUnmarshalErr)
		if errError != nil {
			panic(errError)
		}
	}

}
