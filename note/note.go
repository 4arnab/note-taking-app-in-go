package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() {
	fmt.Println(n.Title, n.Content, n.CreatedAt)
}

func (n Note) SaveToFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", json, 0644)
	// return nil
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return &Note{}, errors.New("invalid input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
