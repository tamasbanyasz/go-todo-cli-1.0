package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type ListOfDays struct {
	daysInSlice []Day
}

type Day struct {
	dayName   string
	taskOfDay []string
}

func CreateDays() []string {
	days := []string{}

	daysOfWeek := []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}

	for _, day := range daysOfWeek {
		days = append(days, day.String())
	}

	return days
}

func (todo *ListOfDays) AppendDaysToListOfDays(days []string) {
	for _, day := range days {
		todo.daysInSlice = append(todo.daysInSlice, Day{dayName: day})
	}
}

func FilesAreExist(dirPath string, days []string) {
	for _, day := range days {

		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error during creating folder:", err)
			return

		} else {
			if f, errPath := os.Stat(dirPath + "/" + day + ".txt"); errPath == nil {
				fmt.Printf("\nFile %s exist", f.Name())

			} else if errors.Is(errPath, os.ErrNotExist) {
				f, errCreate := os.Create(dirPath + "/" + day + ".txt")
				if errCreate != nil {
					fmt.Println(errCreate)
				}
				fmt.Printf("\nFile %s created", f.Name())
				defer f.Close()
			}
		}
		fmt.Printf("\n\n")
	}
}

func (todo *ListOfDays) SaveToFile(dirPath string) {
	for index, day := range todo.daysInSlice {
		emptyString := day.taskOfDay[0]
		err := ioutil.WriteFile(dirPath+"/"+day.dayName+".txt", []byte(emptyString), 0644)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		todo.daysInSlice[index].taskOfDay = nil
	}
}

func (todo *ListOfDays) ReadDaysFromFile(dirPath string, days []string) {
	for index, day := range days {
		readDay, err := ioutil.ReadFile(dirPath + "/" + day + ".txt")
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		}
		if err == nil {
			todo.daysInSlice[index].taskOfDay = append(todo.daysInSlice[index].taskOfDay, string(readDay))
		}
	}
}

func (todo *ListOfDays) AppendTaskToDay(selectedIndex int, task string) {
	for index, day := range todo.daysInSlice {
		if selectedIndex == index {
			fmt.Printf("\nChoosed day %s\n", day)
			todo.daysInSlice[selectedIndex].taskOfDay = nil
			todo.daysInSlice[selectedIndex].taskOfDay = append(todo.daysInSlice[selectedIndex].taskOfDay, task)
		}
	}
}

func main() {

	dirPath := "days_folder"

	var Todo ListOfDays

	createdDays := CreateDays()

	structuredDays := []Day{}

	for _, day := range createdDays {
		structuredDays = append(structuredDays, Day{dayName: day})
	}

	Todo.AppendDaysToListOfDays(createdDays)
	FilesAreExist(dirPath, createdDays)

	for {

		Todo.ReadDaysFromFile(dirPath, createdDays)
		for index, day := range Todo.daysInSlice {
			fmt.Println(index, day.dayName, day.taskOfDay)
		}

		fmt.Println("\nSelect index: ")
		var selectedIndex int
		fmt.Scanln(&selectedIndex)

		fmt.Println("\nTask is: ")
		var task string
		fmt.Scanln(&task)

		Todo.AppendTaskToDay(selectedIndex, task)
		Todo.SaveToFile(dirPath)
	}
}
