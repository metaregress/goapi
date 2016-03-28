package main

import "fmt"

var currentId int

var todoRepo Todos

//create seed data
func init() {
	todoRepo = make(map[int]Todo)
	currentId = 0
	RepoCreateTodo(Todo{Name: "Write presentation?"})
	RepoCreateTodo(Todo{Name: "Host meetup!"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todoRepo {
		if t.Id == id {
			return t
		}
	}
	//return empty todo if none found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todoRepo = append(todoRepo, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todoRepo {
		if t.Id == id {
			todoRepo = append(todoRepo[:i], todoRepo[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find todo with id %d to delete", id)
}
