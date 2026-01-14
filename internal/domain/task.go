package domain

type Task struct {
	Title       string //Заголовок задачи
	Text        string // Описание задачи
	IsDone      bool   // Метка о выполнении
	CreatedAt   string // Время создания
	CompletedAt string // Время завершения
}
