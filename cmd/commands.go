package cmd

import (
	"bufio"
	"endpoint-interview/directory"
	"fmt"
	"os"
	"strings"
)

func ProcessCommandsFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			// fail silently
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "CREATE":
			if len(parts) == 2 {
				directory.CreateDir(parts[1])
			} else {
				fmt.Println("Invalid CREATE command")
			}
		case "MOVE":
			if len(parts) == 3 {
				directory.MoveDir(parts[1], parts[2])
			} else {
				fmt.Println("Invalid MOVE command")
			}
		case "DELETE":
			fmt.Printf("DELETE %s\n", parts[1])
			if len(parts) == 2 {
				directory.DeleteDir(parts[1])
			} else {
				fmt.Println("Invalid DELETE command")
			}
		case "LIST":
			fmt.Println("LIST")
			directory.ListDirs(directory.Root(), "")
		default:
			fmt.Printf("Unknown command: %s\n", parts[0])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
