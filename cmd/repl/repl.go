package repl

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Print(">> ")
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
	fmt.Println("Output ", str)
}
