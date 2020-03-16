package assistservice

type Answers struct {
	Type        string  `json:"type"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	Text        string  `json:"text,omitempty"`
	Choices     Choices `json:"choices,omitempty"`
	Boolean     bool    `json:"boolean,omitempty"`
}

type WebHookRequest struct {
	FormResponse FormResponse `json:"form_response"`
}

type FormResponse struct {
	Answers []Answers `json:"answers"`
}

type Choices struct {
	ID     string   `json:"id"`
	Labels []string `json:"labels"`
}
