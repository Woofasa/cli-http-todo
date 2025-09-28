package usecase

type TaskInput struct {
	Title       string `json:"title"  db:"title"`
	Description string `json:"description"  db:"description"`
}

type UpdateTaskDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Status      *bool   `json:"status,omitempty"`
}

type UserInput struct {
	Name     string `json:"name"  db:"name"`
	Email    string `json:"email"  db:"email"`
	Password string `json:"password"  db:"password"`
}
