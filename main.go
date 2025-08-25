package main

import (
	"fmt"
	"os"

	"github.com/henryhoangduong/go_db_tutorial/cmd/repl"
)

func main() {
	cmd := repl.New()
	cmd.StartREPL()
	fmt.Println("Good bye ...........")
	os.Exit(0)
}
