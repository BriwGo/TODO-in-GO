package storage

import (
	"ToDo/internal/domain"
	errors "ToDo/internal/errs"
)

type Storage struct {
	tasks map[string]domain.Task
}

func NewStorage() *Storage {
	storage := Storage{
		tasks: make(map[string]domain.Task),
	}
	return &storage
}

// метод для добавляения задачи и валидации
func (s *Storage) Add(task domain.Task) {

	s.tasks[task.Title] = task

}

//метод возврата всех задач

func (s *Storage) GetAll() map[string]domain.Task {
	return s.tasks
}

func (s *Storage) Delete(title string) error {

	_, ok := s.tasks[title]
	if !ok {
		return errors.TaskNotFound

	}

	delete(s.tasks, title)
	return nil
}

func (s *Storage) MarkDone(title string) error {

	task, ok := s.tasks[title]
	if !ok {
		return errors.TaskNotFound
	}

	task.Done()

	s.tasks[title] = task
	return nil
}
