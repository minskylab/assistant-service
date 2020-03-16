package main

import "time"

type AutoGenerated struct {
	EventID      string       `json:"event_id"`
	EventType    string       `json:"event_type"`
	FormResponse FormResponse `json:"form_response"`
}

type Fields struct {
	ID                      string    `json:"id"`
	Title                   string    `json:"title"`
	Type                    string    `json:"type"`
	Ref                     string    `json:"ref"`
	AllowMultipleSelections bool      `json:"allow_multiple_selections,omitempty"`
	Choices                 []Choices `json:"choices,omitempty"`
}

type Choices struct {
	ID     string   `json:"id"`
	Labels []string `json:"labels"`
}
type Definition struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Fields []Fields `json:"fields"`
}
type Field struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Ref  string `json:"ref"`
}

type Answers struct {
	Type        string  `json:"type"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	Field       Field   `json:"field"`
	Text        string  `json:"text,omitempty"`
	Choices     Choices `json:"choices,omitempty"`
	Boolean     bool    `json:"boolean,omitempty"`
}
type FormResponse struct {
	FormID      string     `json:"form_id"`
	Token       string     `json:"token"`
	LandedAt    time.Time  `json:"landed_at"`
	SubmittedAt time.Time  `json:"submitted_at"`
	Definition  Definition `json:"definition"`
	Answers     []Answers  `json:"answers"`
}