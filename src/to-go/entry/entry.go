package entry

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type TodoEntry struct {
	ID          string    `csv:"ID"`
	Todo        string    `csv:"Task"`
	IsCompleted bool      `csv:"IsComplete"`
	CreatedAt   time.Time `csv:"CreatedAt"`
	CompletedAt time.Time `csv:"CompletedAt"`
}

func New(todo string) *TodoEntry {
	now := time.Now()
	id := getId(string(todo), now)
	te := &TodoEntry{
		Todo:        todo,
		ID:          id,
		IsCompleted: false,
		CreatedAt:   now,
		CompletedAt: time.Time{},
	}
	return te
}

func (te *TodoEntry) String() string {
	var out bytes.Buffer
	//fmt.Fprintf(&out, "Todo item: %q\nComplete: %t\nCreated At: %v\nCompleted At: %v\n", te.Todo, te.IsCompleted, te.CreatedAt.Format(time.RFC3339), te.CompletedAt.Format(time.RFC3339))
	fmt.Fprintf(&out, "ID: %v\n Todo item: %q\n Complete: %t\n Created At: %v\n Completed At: %v\n", te.ID, te.Todo, te.IsCompleted, te.CreatedAt.Format(time.RFC3339), te.CompletedAt.Format(time.RFC3339))
	return out.String()
}

func (te *TodoEntry) Csv() string {
	var out bytes.Buffer
	fmt.Fprintf(&out, "%v,%v,%t,%v,%v\n", te.ID, te.Todo, te.IsCompleted, te.CreatedAt.Format(time.RFC3339), te.CompletedAt.Format(time.RFC3339))
	return out.String()
}

func (te *TodoEntry) CompleteTodo() {
	if te.IsCompleted == false {
		te.IsCompleted = true
	}
}

func getId(title string, createdAt time.Time) string {
	data := fmt.Sprintf("%s-%d", title, createdAt.UnixNano())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])[:10]
}
