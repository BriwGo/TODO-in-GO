package app

import (
	"ToDo/internal/domain"
	"fmt"

	"github.com/k0kubun/pp"
)

func PrintPrompt() {
	fmt.Print("Enter a Command")
}

func PrintExit() {
	fmt.Print("Todo is Over, GoodBye!")
}

func AddTrue() {
	fmt.Println("Task is Add")
}

func printTasks(s map[string]domain.Task) {
	pp.Println(s)
}

func PrintDone(title string) {
	fmt.Println("Task:", title, "is Done ")
}

func PrintHelp() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func PrintEvents(events []domain.Event) {
	pp.Println("Events:", events)

}
