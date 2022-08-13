package todo

import (
	"errors"
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

	return nil
}
