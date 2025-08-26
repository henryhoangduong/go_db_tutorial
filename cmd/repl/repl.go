package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/henryhoangduong/go_db_tutorial/internal/compiler"
)

type REPL interface {
	StartREPL()
}

type replStruct struct {
}

func New() REPL {
	return &replStruct{}
}

func (r *replStruct) StartREPL() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("DB >> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v \n", err)
			continue
		}
		input = input[:len(input)-1]
		if input == "exit" {
			fmt.Println("Exiting..........")
			break
		}
		r.evaluate(input)
	}
}

func (r *replStruct) evaluate(str string) {
	comp := compiler.NewCompiler(str)
	t := comp.Call()
	fmt.Println("Output ", t)
}
