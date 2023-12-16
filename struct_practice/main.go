package main

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"example.com/struct-practice/note"
)

func main() {

	noteTitle, noteContent, err := getNoteData()

	if err != nil {
		panic(err)
	}

	noteData := note.New(noteTitle, noteContent)

	noteData.Display()

	noteData.Save()
}

func getNoteData() (noteTitle, noteContent string, err error) {
	noteTitle, err = getUserInput("Note title: ")

	if err != nil {
		return "", "", err
	}

	noteContent, err = getUserInput("Note content: ")

	if err != nil {
		return "", "", err
	}

	return noteTitle, noteContent, nil
}

func getUserInput(prompt string) (string, error) {
	print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("Invalid input!")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}
