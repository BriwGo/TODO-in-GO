package domain

import "time"

//структура ивент ЧТо должнна содержать?

type Event struct {
	TimesTamp   time.Time //время когда произошло
	RawInput    string    //что ввел пользователь
	Description string    //текст пользователя

}

func NewEvent(rwinp string, dsrpt string) Event {
	return Event{
		TimesTamp:   time.Now(),
		RawInput:    rwinp,
		Description: dsrpt,
	}
}
