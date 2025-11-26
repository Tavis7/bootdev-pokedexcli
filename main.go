package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		valid := scanner.Scan()
		command := ""
		if valid {
			input := scanner.Text()
			command_line := cleanInput(input)
			if len(command_line) > 0 {
				command = command_line[0]
			}
		} else {
			break
		}
		fmt.Printf("Your command was: %v\n", command)
	}
}
