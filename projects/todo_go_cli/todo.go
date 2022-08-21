package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	// ls := *t // do this to not write (*t) on each line to de-reference the t

	if index <= 0 || index > len(*t) {
		return errors.New("Index out bound")
	}
	(*t)[index-1].CompletedAt = time.Now()
	(*t)[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("Index out bound")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, error := os.ReadFile(filename)

	if error != nil {
		if errors.Is(error, os.ErrNotExist) {
			fmt.Println("No file found")
			return nil
		}
		return error
	}

	if len(file) == 0 {
		// return errors.New("Empty file")
		return nil
	}

	error = json.Unmarshal(file, t)
	if error != nil {
		return error
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, error := json.Marshal(t)
	if error != nil {
		return error
	}

	return os.WriteFile(filename, data, 0644)
}
