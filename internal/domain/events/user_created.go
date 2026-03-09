package events

const (
	UserCreated = "user.created"
)

// UserCreatedEvent TODO: struct of UserCreated consumer body
type UserCreatedEvent struct {
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
	Email  string `json:"email,omitempty"`
}
