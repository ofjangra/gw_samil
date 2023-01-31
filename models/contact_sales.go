package models

type SalesContactMessage struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	WorkEmail   string `json:"work_email" bson:"work_email"`
	Phone       string `json:"phone"`
	MessageText string `json:"message_text" bson:"message_text"`
}
