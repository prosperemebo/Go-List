package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	label      string
	isComplete bool
}

var commands = map[string]func(){
	"t": newTask,
	"c": showCommands,
	"a": showTasks,
	"e": endProgram,
	"p": showTasksPreview,
	"m": markAsComplete,
	"n": markAsNotComplete,
	"r": removeTask,
}

var userTasks = make([]Task, 0)

func main() {
	greet()

	for {
		operation, command := readCommand()

		if operation != nil {
			operation()
		}

		if command == "e" {
			break
		}
	}
}

func greet() {
	fmt.Println("====================")
	fmt.Println("Welcome to GO List")
	fmt.Println("====================")
	fmt.Println("")
	showTip()
}

func readCommand() (func(), string) {
	var userCommand string

	fmt.Print("ENTER YOUR COMMAND: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	userCommand = strings.TrimSpace(text)

	value, validKey := commands[userCommand]

	if !validKey {
		fmt.Printf("[%v] is not a valid command!\n", userCommand)
		fmt.Println("")
		return nil, userCommand
	}

	return value, userCommand
}

func showCommands() {
	fmt.Println("Enter [a] to show all tasks.")
	fmt.Println("Enter [p] to preview tasks.")
	fmt.Println("Enter [t] to add new task.")
	fmt.Println("Enter [m] to mark task as complete.")
	fmt.Println("Enter [n] to mark task as not complete.")
	fmt.Println("Enter [c] to show commands.")
	fmt.Println("Enter [e] to exit program.")
	fmt.Println("")
}

func showTasksPreview() {
	if emptyTasks() {
		return
	}

	fmt.Println("--------------------")
	fmt.Println("To Do")
	fmt.Println("--------------------")

	var completedTasks, unCompletedTasks = sortTasks()

	for index, task := range userTasks {
		if index >= 3 {
			break
		}

		fmt.Printf("[%v] %v\n", index, task.label)
	}

	if len(unCompletedTasks) > 0 {
		fmt.Printf("You have %v more tasks to complete. ", len(unCompletedTasks))
	}

	if len(completedTasks) > 0 {
		fmt.Printf("You have %v complete tasks. ", len(completedTasks))
	}

	fmt.Print("\n")

	fmt.Println("--------------------")

	fmt.Print("\n")

	showTip()
}

func showTasks() {
	if emptyTasks() {
		return
	}

	var completedTasks, unCompletedTasks = sortTasks()

	for index, task := range unCompletedTasks {
		if index == 0 {
			fmt.Println("--------------------")
			fmt.Println("To Do")
			fmt.Println("--------------------")
		}

		fmt.Printf("[%v] %v\n", index, task.label)
	}

	fmt.Println("")

	for index, task := range completedTasks {
		actualIndex := len(unCompletedTasks) + index

		if index == 0 {
			fmt.Println("--------------------")
			fmt.Println("Completed Tasks")
			fmt.Println("--------------------")
		}

		fmt.Printf("[%v] %v\n", actualIndex, task.label)
	}

	fmt.Println("")

}

func newTask() {
	var task string

	fmt.Print("What do you need to do (Enter [x] to cancel)? ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	task = strings.TrimSpace(text)

	if task == "x" {
		return
	}

	var newTask = Task{
		label:      task,
		isComplete: false,
	}

	userTasks = append(userTasks, newTask)

	fmt.Println("New task added.")
	fmt.Println("")
	showTasksPreview()
}

func endProgram() {
	fmt.Println("GoList will shut down.")
	fmt.Println("Goodbye...")
}

func showTip() {
	if len(userTasks) > 0 {
		fmt.Println("Enter [a] to show all tasks.")
	}

	fmt.Println("Enter [c] to show commands.")
	fmt.Println("Enter [t] to add new task.")
}

func sortTasks() ([]Task, []Task) {
	var completed = make([]Task, 0)
	var unCompleted = make([]Task, 0)

	for _, userTask := range userTasks {
		if userTask.isComplete {
			completed = append(completed, userTask)
		} else {
			unCompleted = append(unCompleted, userTask)
		}
	}

	userTasks = append(completed, unCompleted...)

	return completed, unCompleted
}

func markAsComplete() {
	selectedIndex, didExit := validateCheckerAction("Please Enter the task number you want to mark as complete (Enter [x] to cancel): ")

	if didExit {
		return
	}

	var completed = make([]Task, 0)
	var unCompleted = make([]Task, 0)

	for index, userTask := range userTasks {
		if selectedIndex == index {
			userTask = Task{
				label:      userTask.label,
				isComplete: true,
			}
		}

		if userTask.isComplete {
			completed = append(completed, userTask)
		} else {
			unCompleted = append(unCompleted, userTask)
		}
	}

	fmt.Printf("Task [%v] marked as completed!\n", selectedIndex)
	sortTasks()

	userTasks = append(completed, unCompleted...)
}

func markAsNotComplete() {
	selectedIndex, didExit := validateCheckerAction("Please Enter the task number you want to mark as not complete (Enter [x] to cancel): ")

	if didExit {
		return
	}

	var completed = make([]Task, 0)
	var unCompleted = make([]Task, 0)

	for index, userTask := range userTasks {
		if selectedIndex == index {
			userTask = Task{
				label:      userTask.label,
				isComplete: false,
			}
		}

		if userTask.isComplete {
			completed = append(completed, userTask)
		} else {
			unCompleted = append(unCompleted, userTask)
		}
	}

	fmt.Printf("Task [%v] marked as not completed!\n", selectedIndex)
	sortTasks()

	userTasks = append(completed, unCompleted...)
}

func removeTask() {
	selectedIndex, didExit := validateCheckerAction("Please Enter the task number you want to remove (Enter [x] to cancel): ")

	if didExit {
		return
	}

	var completed = make([]Task, 0)
	var unCompleted = make([]Task, 0)

	for index, userTask := range userTasks {
		if selectedIndex == index {
			continue
		}

		if userTask.isComplete {
			completed = append(completed, userTask)
		} else {
			unCompleted = append(unCompleted, userTask)
		}
	}

	fmt.Printf("Task [%v] removed!\n", selectedIndex)
	sortTasks()

	userTasks = append(completed, unCompleted...)
}

func validateCheckerAction(message string) (int, bool) {
	if emptyTasks() {
		return 0, true
	}

	showTasksPreview()
	fmt.Print(message)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	if text == "x" {
		return 0, true
	}

	selectedIndex, err := strconv.ParseInt(strings.TrimSpace(text), 10, 64)

	if err != nil || selectedIndex > int64(len(userTasks)-1) {
		fmt.Printf("Selected task [%v] is not valid\n", selectedIndex)
		fmt.Print("Please select")

		for index, _ := range userTasks {
			if index == len(userTasks)-1 {
				fmt.Printf(" %v.\n", index)
			} else {
				fmt.Printf(" %v or", index)
			}
		}

		validateCheckerAction(message)
	}

	return int(selectedIndex), false
}

func emptyTasks() bool {
	if len(userTasks) == 0 {
		fmt.Println("You don't have any tasks!")
		showTip()
		fmt.Println("")
		return true
	}

	return false
}
