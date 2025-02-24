package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	//"github.com/jackjf28/to-go/entry"
	"io/fs"
	//"path/filepath"
	//"io"
	//"log"
	"os"
	"strings"
)

const DATA_DIR = "./data"

var ()

func checkForDataDir() {
	dirExists, err := exists(DATA_DIR)
	check(err)
	if !dirExists {
		os.Mkdir(DATA_DIR, 0755)
	}
}

func formatEntry(s string) string {
	trimmed := strings.TrimSpace(s)
	result := trimmed + "\n"
	return result
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func openTodoFile(path string) (*os.File, error) {
	fileExists, err := exists(path)
	check(err)
	if fileExists {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0644)
		return f, err
	} else {
		//f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
		f, err := os.Create(path)
		return f, err
	}
}

func handleNewEntry(entry string) {
	fmt.Println("Inserting entry: ", entry)

	f, err := openTodoFile("./data/todos.csv")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	line := formatEntry(entry)
	n, err := w.Write([]byte(line))
	w.Flush()
	fmt.Printf("Wrote %d bytes.\n", n)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	entryCmd := flag.NewFlagSet("entry", flag.ExitOnError)
	entryNew := entryCmd.String("new", "", "create a new entry")

	if len(os.Args) < 2 {
		fmt.Println("'new' command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "entry":
		checkForDataDir()
		entryCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'entry'")
		fmt.Println("  new:", *entryNew)
		if *entryNew != "" {
			handleNewEntry(*entryNew)
		}
	default:
		fmt.Println("expected 'entry' subcommands")
		os.Exit(1)
	}
}

//r := strings.NewReader("hello, reader!")
//b := make([]byte, 8)
//for {
//	n, err := r.Read(b)
//	if err == io.EOF {
//		break
//	}
//}
