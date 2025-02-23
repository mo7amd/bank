package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Account struct {
	Id        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Amount      float64   `json:"amount"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Event struct {
	EventType string      `json:"event_type"`
	Data      interface{} `json:"data"`
	Metadata  struct {
		Timestamp time.Time `json:"timestamp"`
		TraceID   string    `json:"trace_id"`
	} `json:"metadata"`
}
