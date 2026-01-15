package app

import (
	"ToDo/internal/domain"
	"ToDo/internal/storage"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	store *storage.Storage
}

func (s *Service) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		PrintPrompt()

		ok := scanner.Scan()
		if !ok {
			return
		}

		inputString := scanner.Text()

		result, err := s.ProcessComand(inputString)
		if err != nil {
			fmt.Println(err)
		}

		if result == "e" {
			PrintExit()
			return
		}

	}
}

func (s *Service) ProcessComand(inputstr string) (string, error) {
	fields := strings.Fields(inputstr)

	if len(fields) == 0 {
		return "", EmptyInput
	}

	cmd := fields[0]

	if cmd == "Exit" || cmd == "exit" {
		Exit := "e"
		return Exit, nil

	}

	if cmd == "Add" || cmd == "add" {

		add, err := s.Add(fields)
		if err != nil {
			fmt.Println(err)
		}

		AddTrue()

		return add, nil
	}

	if cmd == "List" || cmd == "list" {


	}

	return "", nil
}

func (s *Service) Add(fields []string) (string, error) {

	if len(fields) < 3 {
		return "", WrongArgs
	}

	title := fields[1]

	taskText := ""

	for i := 2; i < len(fields); i++ {

		taskText += fields[i]

		if i != len(fields)-1 {
			taskText += " "
		}

		task := domain.NewTask(title, taskText)
		s.store.Add(task)

	}

	return "Great", nil
}


func 