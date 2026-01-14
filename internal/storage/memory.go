package storage

import "ToDo/internal/domain"

//структура сторадж ЧТо должнна содержать?
type Storage struct {
	tasks  map[string]domain.Task
	events []domain.Event
}

//конструктор новый сторадж для валидации
func NewStorage() {

}

//метод для добавляения задачи и валидации
func (p *Storage) Add() {

}

//метод возврата всех задач

func (p Storage) GetAll() {

}

//метод удаления задачи
func (p *Storage) Delete() {

}

// метод для того чтобы ставить задаче статус выполнено
func (p *Storage) MarkDone() {

}

// метод для возврата событий
func (p Storage) GetEvents() {

}
