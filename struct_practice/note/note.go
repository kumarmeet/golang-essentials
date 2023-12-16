package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct tags (meta deta)
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// method
func (n Note) Display() {
	fmt.Println(n.Title, n.Content, n.CreatedAt)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "-")
	fileName = strings.ToLower(fileName) + ".json"

	// Marshal function will only extract and convert struct data that made publicly available (eg. capitalize struct fields)
	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// creation or constructor function
func New(title, content string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
