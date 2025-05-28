package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/4arnab/note-taking-app/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	userNote.Display()
	userNote.SaveToFile()

	err = userNote.SaveToFile()

	if err != nil {
		fmt.Println(err)
		fmt.Print("FAILED TO SAVE FILE")
		panic("EXIT")
	}

	fmt.Print("SUCCESS TO SAVE FILE")
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
