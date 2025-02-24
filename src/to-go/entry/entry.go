package entry

import (
	"bytes"
	"fmt"
)

type Todo string

type TodoEntry struct {
	Todo
	IsCompleted bool
}

func New(todo Todo, isComplete bool) *TodoEntry {
	te := &TodoEntry{
		Todo:        todo,
		IsCompleted: isComplete,
	}
	return te
}

func (te *TodoEntry) String() string {
	var out bytes.Buffer
	fmt.Fprintf(&out, "Todo item: %q\nIs Complete: %t\n", te.Todo, te.IsCompleted)
	return out.String()
}

func (te *TodoEntry) Csv() {}
