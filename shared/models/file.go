package models

type File struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Size     int64             `json:"size"`
	Metadata map[string]string `json:"metadata"`
}
