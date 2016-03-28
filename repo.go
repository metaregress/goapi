package main

import "fmt"

var nextId int

var todoRepo Todos

//create seed data
func init() {
	todoRepo = make(map[int]Todo)
	nextId = 0
	RepoCreateTodo(Todo{Name: "Make more todos."})
}

func RepoFindTodo(id int) Todo {
	todo, ok := todoRepo[id]
	if !ok {
		//return empty todo if none found
		return Todo{}
	}
	return todo
}

func RepoCreateTodo(t Todo) Todo {
	t.Id = nextId
	todoRepo[nextId] = t
	nextId += 1
	return t
}

func RepoDestroyTodo(id int) error {
	_, ok := todoRepo[id]
	if !ok {
		return fmt.Errorf("Could not find todo with id %d to delete", id)
	}
	delete(todoRepo, id)
	return nil
}
