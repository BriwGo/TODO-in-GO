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
	store  *storage.Storage
	events []domain.Event
}

func NewService(st *storage.Storage) *Service {
	return &Service{
		store:  st,
		events: make([]domain.Event, 0),
	}
}

func (s *Service) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		PrintPrompt()

		if !scanner.Scan() {
			PrintExit()
			return
		}

		inputString := scanner.Text()

		shouldExit, err := s.ProcessCommand(inputString)

		var status string
		if err != nil {
			fmt.Println("Ошибка:", err)
			status = fmt.Sprintf("Error: %v", err)
		} else {
			status = "Success"
		}

		event := domain.NewEvent(inputString, status)
		s.events = append(s.events, event)

		if shouldExit {
			PrintExit()
			return
		}
	}
}
func (s *Service) ProcessCommand(inputStr string) (shouldExit bool, err error) {
	fields := strings.Fields(inputStr)

	if len(fields) == 0 {
		return false, EmptyInput
	}

	cmd := fields[0]

	switch cmd {

	case "Events", "events":
		err := s.Events(fields)
		if err != nil {
			return false, err
		}
		return false, nil

	case "Help", "help":
		err := s.Help(fields)
		if err != nil {
			return false, err
		}
		return false, nil

	case "Exit", "exit":
		return true, nil

	case "Add", "add":
		err := s.Add(fields)
		if err != nil {
			return false, err
		}
		AddTrue()
		return false, nil

	case "List", "list":
		err := s.List(fields)
		if err != nil {
			return false, err
		}
		return false, nil

	case "Done", "done":
		err := s.Done(fields)
		if err != nil {
			return false, err
		}
		return false, nil

	case "Del", "del":
		err := s.Delete(fields)
		if err != nil {
			return false, err
		}
		return false, nil

	default:
		return false, fmt.Errorf("unknown command: %s", cmd)
	}
}

func (s *Service) Events(fields []string) error {
	if len(fields) != 1 {
		return WrongArgs
	}
	PrintEvents(s.events)
	return nil
}

func (s *Service) Help(fields []string) error {

	if len(fields) != 1 {
		return WrongArgs
	}
	PrintHelp()
	return nil

}
func (s *Service) Add(fields []string) error {
	if len(fields) < 3 {
		return WrongArgs
	}

	title := fields[1]
	taskText := strings.Join(fields[2:], " ")

	task := domain.NewTask(title, taskText)
	s.store.Add(task)

	return nil
}

func (s *Service) List(fields []string) error {
	if len(fields) != 1 {
		return WrongArgs
	}

	tasks := s.store.GetAll()
	printTasks(tasks)
	return nil
}

func (s *Service) Done(fields []string) error {
	if len(fields) != 2 {
		return WrongArgs
	}

	title := fields[1]
	err := s.store.MarkDone(title)
	if err != nil {
		return err
	}

	PrintDone(title)
	return nil

}

func (s *Service) Delete(fields []string) error {
	if len(fields) != 2 {
		return WrongArgs
	}
	title := fields[1]
	err := s.store.Delete(title)
	if err != nil {
		return err
	}
	return nil

}
