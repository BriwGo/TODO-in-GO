package app

import (
	"ToDo/internal/storage"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	store *storage.Storage
}

// тут паристится строка и передается в командс го
// Берет строуку  по принципу : команда,аргумент 1,аргумент 2
func Parser() {

	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		fmt.Println("Input Error")
	}

	text := scanner.Text()
	fields := strings.Fields(text)

	if len(fields) == 0 {
		fmt.Println("Lenght to low")
	}

	cmd := fields[0]

	if cmd == "leave" || cmd == "Leave" {
		fmt.Println("OK")
		return
	}

	if cmd == "Add" || cmd == "add" {

	} else if cmd == "Delete" || cmd == "delete" {

	}

}

// метод который должен парсить, определять первое слово(пользовательский ввод)
// в зависимости от команды указывает место хранилищу

func (s *Service) ProcessComand(input string) {

}
