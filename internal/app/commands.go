package app

import "ToDo/internal/storage"

type Service struct {
	store *storage.Storage
}

//тут паристится строка и передается в командс го
//Берет строуку  по принципу : команда,аргумент 1,аргумент 2
func Parser(NewStroke string) {

}

// метод который должен парсить, определять первое слово(пользовательский ввод)
// в зависимости от команды указывает место хранилищу

func (s *Service) ProcessComand(input string) {

}
