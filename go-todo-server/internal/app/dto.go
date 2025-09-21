package app

type TaskInput struct {
	Title       string `json:"title"  db:"title"`
	Description string `json:"description"  db:"description"`
}

type UpdateDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}
