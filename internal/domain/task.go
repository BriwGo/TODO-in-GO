package domain

import "time"

type Task struct {
	Title       string     //Заголовок задачи
	Text        string     // Описание задачи
	IsDone      bool       // Метка о выполнении
	CreatedAt   time.Time  // Время создания
	CompletedAt *time.Time // Время завершения
}

func NewTask(title, text string) Task {
	task := Task{
		Title:       title,
		Text:        text,
		IsDone:      false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	return task
}

func (t *Task) Done() {

	t.IsDone = true
	now := time.Now()
	t.CompletedAt = &now
}
