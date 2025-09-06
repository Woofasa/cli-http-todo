package app

type TaskInput struct {
	Title       string `json:"title"  db:"title"`
	Description string `json:"description"  db:"description"`
}
