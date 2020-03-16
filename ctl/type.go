package main

import "time"

type Answers struct {
	Type        string  `json:"type"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	Text        string  `json:"text,omitempty"`
	Choices     Choices `json:"choices,omitempty"`
	Boolean     bool    `json:"boolean,omitempty"`
}
type FormResponse struct {
	FormID      string    `json:"form_id"`
	Token       string    `json:"token"`
	LandedAt    time.Time `json:"landed_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	Answers     []Answers `json:"answers"`
}

type Choices struct {
	ID     string   `json:"id"`
	Labels []string `json:"labels"`
}
